package main

import (
	"fmt"

	"github.com/golang/glog"
	v1 "github.com/openshift/api/image/v1"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
)

func getNonVerifiedTags(acceptedTags []*v1.TagReference) (current, previous *v1.TagReference) {
	// We need at least 2 accepted releases to compare
	if len(acceptedTags) < 2 {
		return nil, nil
	}
	// get latest accepted tag that has not been bugzilla verified
	for index, tag := range acceptedTags {
		if anno, ok := tag.Annotations[releaseAnnotationBugsVerified]; ok && anno == "true" {
			// if we've already checked the latest tag, return
			if index == 0 {
				return nil, nil
			}
			previous = tag
			current = acceptedTags[index-1]
			break
		}
	}
	if previous == nil || current == nil {
		// handle case where no tags have been verified yet; start with oldest tags
		previous = acceptedTags[len(acceptedTags)-1]
		current = acceptedTags[len(acceptedTags)-2]
	}
	return current, previous
}

// syncBugzilla checks whether fixed bugs in a release had their
// PR reviewed and approved by the QA contact for the bug
func (c *Controller) syncBugzilla(key queueKey) error {
	defer func() {
		err := recover()
		panic(err)
	}()

	release, err := c.loadReleaseForSync(key.namespace, key.name)
	if err != nil || release == nil {
		return err
	}

	glog.V(6).Infof("checking if %v (%s) has verifyBugs set", key, release.Config.Name)

	// check if verifyBugs publish step is enabled for release
	var verifyBugs *PublishVerifyBugs
	for _, publishType := range release.Config.Publish {
		if publishType.Disabled {
			continue
		}
		switch {
		case publishType.VerifyBugs != nil:
			verifyBugs = publishType.VerifyBugs
		}
		if verifyBugs != nil {
			break
		}
	}
	if verifyBugs == nil {
		glog.V(6).Infof("%v (%s) does not have verifyBugs set", key, release.Config.Name)
		return nil
	}

	glog.V(4).Infof("Verifying fixed bugs in %s", release.Config.Name)

	// get accepted tags
	acceptedTags := sortedRawReleaseTags(release, releasePhaseAccepted)
	tag, prevTag := getNonVerifiedTags(acceptedTags)
	if tag == nil || prevTag == nil {
		glog.V(6).Infof("bugzilla: All accepted tags for %s have already been verified", release.Config.Name)
		return nil
	}

	bugs, err := c.releaseInfo.Bugs(prevTag.Name, tag.Name)
	if err != nil {
		return fmt.Errorf("Unable to generate changelog from %s to %s: %v", prevTag.Name, tag.Name, err)
	}
	var errs []error
	if errs := append(errs, c.bugzillaVerifier.VerifyBugs(bugs)...); len(errs) != 0 {
		return utilerrors.NewAggregate(errs)
	}

	target := release.Target.DeepCopy()
	tagToBeUpdated := findTagReference(target, tag.Name)
	if tagToBeUpdated == nil {
		return fmt.Errorf("release %s no longer exists, cannot set annotation %s=true", tag.Name, releaseAnnotationBugsVerified)
	}
	if tagToBeUpdated.Annotations == nil {
		tagToBeUpdated.Annotations = make(map[string]string)
	}
	tagToBeUpdated.Annotations[releaseAnnotationBugsVerified] = "true"
	if _, err := c.imageClient.ImageStreams(target.Namespace).Update(target); err != nil {
		return err
	}

	return nil
}
