#### Synopsis
---
Create a Go tool that scans an AWS account for unused or underutilized resources and either reports them or automates their cleanup. The tool would use AWS APIs to identify things like: EC2 instances that are stopped or have very low utilization, orphaned EBS volumes or snapshots, unused Elastic IPs, old Lambda functions or versions that are no longer invoked, or even idle EKS clusters or node groups. It could then take action (with confirmation) to delete those resources or tag them for review. For safety, start in “report mode” – listing candidates for cleanup – then add an option to prune. Essentially, this is a cost-savings and housekeeping utility (inspired by tools like **aws-nuke**, which can delete all resources in an account, though your tool would be more selective).  

#### Why I am making this project
---
because i think its cool 
