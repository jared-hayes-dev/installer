package image

import (
	"encoding/json"
	"io"
	"os"

	"github.com/openshift/assisted-image-service/pkg/isoeditor"
	"github.com/openshift/installer/pkg/agent/isosource"
	"github.com/openshift/installer/pkg/asset"
)

const (
	agentISOFilename = "output/agent.iso"
)

// ISO is an asset that generates the bootable image used to install clusters.
type ISO struct {
	File *asset.File
}

var _ asset.WritableAsset = (*ISO)(nil)

// Dependencies returns the assets on which the Bootstrap asset depends.
func (a *ISO) Dependencies() []asset.Asset {
	return []asset.Asset{
		&Ignition{},
	}
}

// Generate generates the image file for to ISO asset.
func (a *ISO) Generate(dependencies asset.Parents) error {
	ignition := &Ignition{}
	dependencies.Get(ignition)

	// TODO: Use BaseIso asset once it is available
	baseImage, err := isosource.EnsureIso()
	if err != nil {
		return err
	}

	ignitionByte, err := json.Marshal(ignition.Config)
	if err != nil {
		return err
	}

	ignitionContent := &isoeditor.IgnitionContent{Config: ignitionByte}
	custom, err := isoeditor.NewRHCOSStreamReader(baseImage, ignitionContent, nil)
	if err != nil {
		return err
	}
	defer custom.Close()

	output, err := os.Create(agentISOFilename)
	if err != nil {
		return err
	}
	defer output.Close()

	_, err = io.Copy(output, custom)
	return err
}

// Name returns the human-friendly name of the asset.
func (a *ISO) Name() string {
	return "Agent Installer ISO"
}

// Load returns the ISO from disk.
func (a *ISO) Load(f asset.FileFetcher) (bool, error) {
	// The ISO will not be needed by another asset so load is noop.
	// This is implemented because it is required by WritableAsset
	return false, nil
}

// Files returns the files generated by the asset.
func (a *ISO) Files() []*asset.File {
	// Return empty array because File will never be loaded.
	return []*asset.File{}
}
