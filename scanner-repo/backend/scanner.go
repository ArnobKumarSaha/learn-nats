/*
Copyright AppsCode Inc. and Contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package backend

import (
	_ "gocloud.dev/blob/s3blob"
	"gomodules.xyz/blobfs"
)

func NewBlobFS() blobfs.Interface {
	// Get or set the AccessKey & SecretKey envs
	storeURL := "s3://scanner-reports?endpoint=https://us-east-1.linodeobjects.com&region=US"
	return blobfs.New(storeURL)
}

func DownloadReport(fs blobfs.Interface, img string) ([]byte, error) {
	// Utilize ParseReference(), & fs.ReadFile /repo/digest/report.json
	return nil, nil
}

func DownloadSummary(fs blobfs.Interface, img string) ([]byte, error) {
	// Utilize ParseReference(), & ReadFile /repo/digest/summary.json
	return nil, nil
}

func ExistsReport(fs blobfs.Interface, img string) (bool, error) {
	// Utilize ParseReference(), & fs.Exists /repo/digest/summary.json
	return false, nil
}

func UploadReport(fs blobfs.Interface, img string) error {
	// report, reportBytes, err := scan(img)
	// repo, digest, err := ParseReference(img)

	// Use the report to calculate the summary.Results
	// fs.WriteFile report.json & summary.json
	return nil
}

// trivy image ubuntu --security-checks vuln --format json --quiet
func scan(img string) ([]byte, error) {
	// Run the `trivy` command, unmarshal it to a Report
	return nil, nil
}

func ParseReference(img string) (repo string, digest string, err error) {
	// Parse the img reference.
	// repo = ref.Context()
	// digest = either crane it, or set the value of ref.Identifier
	return
}
