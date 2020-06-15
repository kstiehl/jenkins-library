// Code generated by piper's step-generator. DO NOT EDIT.

package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/SAP/jenkins-library/pkg/config"
	"github.com/SAP/jenkins-library/pkg/log"
	"github.com/SAP/jenkins-library/pkg/telemetry"
	"github.com/spf13/cobra"
)

type whitesourceExecuteScanOptions struct {
	BuildDescriptorFile                  string `json:"buildDescriptorFile,omitempty"`
	DefaultVersioningModel               string `json:"defaultVersioningModel,omitempty"`
	CreateProductFromPipeline            bool   `json:"createProductFromPipeline,omitempty"`
	SecurityVulnerabilities              bool   `json:"securityVulnerabilities,omitempty"`
	Timeout                              string `json:"timeout,omitempty"`
	AgentDownloadURL                     string `json:"agentDownloadUrl,omitempty"`
	ConfigFilePath                       string `json:"configFilePath,omitempty"`
	VulnerabilityReportFileName          string `json:"vulnerabilityReportFileName,omitempty"`
	ParallelLimit                        string `json:"parallelLimit,omitempty"`
	Reporting                            bool   `json:"reporting,omitempty"`
	ServiceURL                           string `json:"serviceUrl,omitempty"`
	BuildDescriptorExcludeList           string `json:"buildDescriptorExcludeList,omitempty"`
	OrgToken                             string `json:"orgToken,omitempty"`
	UserToken                            string `json:"userToken,omitempty"`
	LicensingVulnerabilities             bool   `json:"licensingVulnerabilities,omitempty"`
	AgentFileName                        string `json:"agentFileName,omitempty"`
	EmailAddressesOfInitialProductAdmins string `json:"emailAddressesOfInitialProductAdmins,omitempty"`
	ProductVersion                       string `json:"productVersion,omitempty"`
	JreDownloadURL                       string `json:"jreDownloadUrl,omitempty"`
	ProductName                          string `json:"productName,omitempty"`
	Verbose                              bool   `json:"verbose,omitempty"`
	ProjectName                          string `json:"projectName,omitempty"`
	ProjectToken                         string `json:"projectToken,omitempty"`
	VulnerabilityReportTitle             string `json:"vulnerabilityReportTitle,omitempty"`
	InstallCommand                       string `json:"installCommand,omitempty"`
	ScanType                             string `json:"scanType,omitempty"`
	CvssSeverityLimit                    string `json:"cvssSeverityLimit,omitempty"`
	ProductToken                         string `json:"productToken,omitempty"`
	AgentParameters                      string `json:"agentParameters,omitempty"`
}

// WhitesourceExecuteScanCommand BETA
func WhitesourceExecuteScanCommand() *cobra.Command {
	const STEP_NAME = "whitesourceExecuteScan"

	metadata := whitesourceExecuteScanMetadata()
	var stepConfig whitesourceExecuteScanOptions
	var startTime time.Time

	var createWhitesourceExecuteScanCmd = &cobra.Command{
		Use:   STEP_NAME,
		Short: "BETA",
		Long: `BETA
With this step [WhiteSource](https://www.whitesourcesoftware.com) security and license compliance scans can be executed and assessed.
WhiteSource is a Software as a Service offering based on a so called unified agent that locally determines the dependency
tree of a node.js, Java, Python, Ruby, or Scala based solution and sends it to the WhiteSource server for a policy based license compliance
check and additional Free and Open Source Software Publicly Known Vulnerabilities detection.
!!! note "Docker Images"
    The underlying Docker images are public and specific to the solution's programming language(s) and therefore may have to be exchanged
    to fit to and support the relevant scenario. The default Python environment used is i.e. Python 3 based.
!!! warn "Restrictions"
    Currently the step does contain hardened scan configurations for ` + "`" + `scanType` + "`" + ` ` + "`" + `'pip'` + "`" + ` and ` + "`" + `'go'` + "`" + `. Other environments are still being elaborated,
    so please thoroughly check your results and do not take them for granted by default.
    Also not all environments have been thoroughly tested already therefore you might need to tweak around with the default containers used or
    create your own ones to adequately support your scenario. To do so please modify ` + "`" + `dockerImage` + "`" + ` and ` + "`" + `dockerWorkspace` + "`" + ` parameters.
    The step expects an environment containing the programming language related compiler/interpreter as well as the related build tool. For a list
    of the supported build tools per environment please refer to the [WhiteSource Unified Agent Documentation](https://whitesource.atlassian.net/wiki/spaces/WD/pages/33718339/Unified+Agent).`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			startTime = time.Now()
			log.SetStepName(STEP_NAME)
			log.SetVerbose(GeneralConfig.Verbose)

			path, _ := os.Getwd()
			fatalHook := &log.FatalHook{CorrelationID: GeneralConfig.CorrelationID, Path: path}
			log.RegisterHook(fatalHook)

			err := PrepareConfig(cmd, &metadata, STEP_NAME, &stepConfig, config.OpenPiperFile)
			if err != nil {
				return err
			}

			if len(GeneralConfig.HookConfig.SentryConfig.Dsn) > 0 {
				sentryHook := log.NewSentryHook(GeneralConfig.HookConfig.SentryConfig.Dsn, GeneralConfig.CorrelationID)
				log.RegisterHook(&sentryHook)
			}

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			telemetryData := telemetry.CustomData{}
			telemetryData.ErrorCode = "1"
			handler := func() {
				telemetryData.Duration = fmt.Sprintf("%v", time.Since(startTime).Milliseconds())
				telemetry.Send(&telemetryData)
			}
			log.DeferExitHandler(handler)
			defer handler()
			telemetry.Initialize(GeneralConfig.NoTelemetry, STEP_NAME)
			whitesourceExecuteScan(stepConfig, &telemetryData)
			telemetryData.ErrorCode = "0"
			log.Entry().Info("SUCCESS")
		},
	}

	addWhitesourceExecuteScanFlags(createWhitesourceExecuteScanCmd, &stepConfig)
	return createWhitesourceExecuteScanCmd
}

func addWhitesourceExecuteScanFlags(cmd *cobra.Command, stepConfig *whitesourceExecuteScanOptions) {
	cmd.Flags().StringVar(&stepConfig.BuildDescriptorFile, "buildDescriptorFile", `./Gopkg.toml`, "Explicit path to the build descriptor file.")
	cmd.Flags().StringVar(&stepConfig.DefaultVersioningModel, "defaultVersioningModel", `major`, "The default project versioning model used in case `projectVersion` parameter is empty for creating the version based on the build descriptor version to report results in Whitesource, can be one of `'major'`, `'major-minor'`, `'semantic'`, `'full'`")
	cmd.Flags().BoolVar(&stepConfig.CreateProductFromPipeline, "createProductFromPipeline", true, "Whether to create the related WhiteSource product on the fly based on the supplied pipeline configuration.")
	cmd.Flags().BoolVar(&stepConfig.SecurityVulnerabilities, "securityVulnerabilities", true, "Whether security compliance is considered and reported as part of the assessment.")
	cmd.Flags().StringVar(&stepConfig.Timeout, "timeout", `0`, "Timeout in seconds until a HTTP call is forcefully terminated.")
	cmd.Flags().StringVar(&stepConfig.AgentDownloadURL, "agentDownloadUrl", `https://github.com/unified-agent-distribution/releases/latest/download/${config.agentFileName}`, "URL used to download the latest version of the WhiteSource Unified Agent.")
	cmd.Flags().StringVar(&stepConfig.ConfigFilePath, "configFilePath", `./wss-unified-agent.config`, "Explicit path to the WhiteSource Unified Agent configuration file.")
	cmd.Flags().StringVar(&stepConfig.VulnerabilityReportFileName, "vulnerabilityReportFileName", `piper_whitesource_vulnerability_report`, "Name of the file the vulnerability report is written to.")
	cmd.Flags().StringVar(&stepConfig.ParallelLimit, "parallelLimit", `15`, "Limit of parallel jobs being run at once in case of `scanType: 'mta'` based scenarios, defaults to `15`.")
	cmd.Flags().BoolVar(&stepConfig.Reporting, "reporting", true, "Whether assessment is being done at all, defaults to `true`.")
	cmd.Flags().StringVar(&stepConfig.ServiceURL, "serviceUrl", `https://saas.whitesourcesoftware.com/api`, "URL to the WhiteSource server API used for communication, defaults to `https://saas.whitesourcesoftware.com/api`.")
	cmd.Flags().StringVar(&stepConfig.BuildDescriptorExcludeList, "buildDescriptorExcludeList", `[]`, "List of build descriptors and therefore modules to exclude from the scan and assessment activities.")
	cmd.Flags().StringVar(&stepConfig.OrgToken, "orgToken", os.Getenv("PIPER_orgToken"), "WhiteSource token identifying your organization.")
	cmd.Flags().StringVar(&stepConfig.UserToken, "userToken", os.Getenv("PIPER_userToken"), "WhiteSource token identifying the user executing the scan")
	cmd.Flags().BoolVar(&stepConfig.LicensingVulnerabilities, "licensingVulnerabilities", true, "Whether license compliance is considered and reported as part of the assessment.")
	cmd.Flags().StringVar(&stepConfig.AgentFileName, "agentFileName", `wss-unified-agent.jar`, "Locally used name for the Unified Agent jar file after download.")
	cmd.Flags().StringVar(&stepConfig.EmailAddressesOfInitialProductAdmins, "emailAddressesOfInitialProductAdmins", `[]`, "The list of email addresses to assign as product admins for newly created WhiteSource products.")
	cmd.Flags().StringVar(&stepConfig.ProductVersion, "productVersion", os.Getenv("PIPER_productVersion"), "Version of the WhiteSource product to be created and used for results aggregation, usually determined automatically.")
	cmd.Flags().StringVar(&stepConfig.JreDownloadURL, "jreDownloadUrl", os.Getenv("PIPER_jreDownloadUrl"), "URL used for downloading the Java Runtime Environment (JRE) required to run the WhiteSource Unified Agent.")
	cmd.Flags().StringVar(&stepConfig.ProductName, "productName", os.Getenv("PIPER_productName"), "Name of the WhiteSource product to be created and used for results aggregation.")
	cmd.Flags().BoolVar(&stepConfig.Verbose, "verbose", false, "Whether verbose output should be produced.")
	cmd.Flags().StringVar(&stepConfig.ProjectName, "projectName", `{{list .GroupID .ArtifactID | join "-" | trimAll "-"}}`, "The project used for reporting results in Whitesource")
	cmd.Flags().StringVar(&stepConfig.ProjectToken, "projectToken", os.Getenv("PIPER_projectToken"), "Project token to execute scan on")
	cmd.Flags().StringVar(&stepConfig.VulnerabilityReportTitle, "vulnerabilityReportTitle", `WhiteSource Security Vulnerability Report`, "Title of vulnerability report written during the assessment phase.")
	cmd.Flags().StringVar(&stepConfig.InstallCommand, "installCommand", os.Getenv("PIPER_installCommand"), "Install command that can be used to populate the default docker image for some scenarios.")
	cmd.Flags().StringVar(&stepConfig.ScanType, "scanType", os.Getenv("PIPER_scanType"), "Type of development stack used to implement the solution.")
	cmd.Flags().StringVar(&stepConfig.CvssSeverityLimit, "cvssSeverityLimit", `7`, "Limit of tollerable CVSS v3 score upon assessment and in consequence fails the build, defaults to  `-1`.")
	cmd.Flags().StringVar(&stepConfig.ProductToken, "productToken", os.Getenv("PIPER_productToken"), "Token of the WhiteSource product to be created and used for results aggregation, usually determined automatically.")
	cmd.Flags().StringVar(&stepConfig.AgentParameters, "agentParameters", ``, "Additional parameters passed to the Unified Agent command line.")

	cmd.MarkFlagRequired("orgToken")
	cmd.MarkFlagRequired("userToken")
	cmd.MarkFlagRequired("productName")
}

// retrieve step metadata
func whitesourceExecuteScanMetadata() config.StepData {
	var theMetaData = config.StepData{
		Metadata: config.StepMetadata{
			Name:    "whitesourceExecuteScan",
			Aliases: []config.Alias{},
		},
		Spec: config.StepSpec{
			Inputs: config.StepInputs{
				Parameters: []config.StepParameters{
					{
						Name:        "buildDescriptorFile",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "buildDescriptorFile",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "buildDescriptorFile",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "buildDescriptorFile",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "buildDescriptorFile",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "buildDescriptorFile",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "buildDescriptorFile",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "defaultVersioningModel",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "createProductFromPipeline",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "bool",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "securityVulnerabilities",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "bool",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "timeout",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "agentDownloadUrl",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "configFilePath",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "vulnerabilityReportFileName",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "parallelLimit",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "reporting",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "bool",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "serviceUrl",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "buildDescriptorExcludeList",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "orgToken",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "userToken",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "licensingVulnerabilities",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "bool",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "agentFileName",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "emailAddressesOfInitialProductAdmins",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "productVersion",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "jreDownloadUrl",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "productName",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "verbose",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:        "bool",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "projectName",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{{Name: "whitesourceProjectName"}},
					},
					{
						Name:        "projectToken",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "vulnerabilityReportTitle",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "installCommand",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "scanType",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "cvssSeverityLimit",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "productToken",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GENERAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
					{
						Name:        "agentParameters",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
					},
				},
			},
		},
	}
	return theMetaData
}
