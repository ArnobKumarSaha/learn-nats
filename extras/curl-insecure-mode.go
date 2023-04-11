func (r *WorkloadReconciler) getVulnerabilityDBLastUpdateTime() (trivy.Time, error) {
	tokenData, err := fs.ReadFile(os.DirFS("/var/run/secrets/kubernetes.io/serviceaccount"), "token")
	if err != nil {
		return trivy.Time{}, err
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	url := "https://scanner.kubeops.svc/var/data/files/trivy/metadata.json"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return trivy.Time{}, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", string(tokenData)))

	res, err := client.Do(req)
	if err != nil {
		return trivy.Time{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return trivy.Time{}, err
	}

	var ver trivy.VulnerabilityDBStruct
	err = json.Unmarshal(body, &ver)
	if err != nil {
		return trivy.Time{}, err
	}
	return ver.UpdatedAt, nil
}
