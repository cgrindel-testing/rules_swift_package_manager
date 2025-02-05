package swift

import (
	"path"
	"strings"

	"github.com/bazelbuild/bazel-gazelle/label"
	"github.com/cgrindel/rules_swift_package_manager/gazelle/internal/swiftpkg"
)

const modulemapLabelNameSuffix = "_modulemap"

// BazelLabelFromTarget creates a Bazel label from a Swift target.
// The logic in this function must stay in sync with
// pkginfo_targets.bazel_label_name_from_parts() in the Starlark code.
func BazelLabelFromTarget(repoName string, target *swiftpkg.Target) *label.Label {
	var name string
	basename := path.Base(target.Path)
	dirname := path.Dir(target.Path)
	if basename == target.Name && dirname != "." {
		name = target.Path
	} else {
		name = path.Join(target.Path, target.Name)
	}
	name = strings.ReplaceAll(name, "/", "_")
	lbl := label.New(repoName, "", name)
	return &lbl
}

// ModulemapBazelLabelFromTargetLabel creates a Bazel label for a modulemap target from the
// corresponding target label.
func ModulemapBazelLabelFromTargetLabel(lbl *label.Label) *label.Label {
	mml := label.New(lbl.Repo, lbl.Pkg, lbl.Name+modulemapLabelNameSuffix)
	return &mml
}
