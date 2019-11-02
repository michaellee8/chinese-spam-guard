import { WebhookPayloadIssuesIssue } from "@octokit/webhooks";

interface SpamScanner {
  scanText(text: string): boolean;
  scanIssue(issue: WebhookPayloadIssuesIssue): boolean;
}

export { SpamScanner };
