import { SpamScanner } from "./spam-scanner";
import { WebhookPayloadIssuesIssue } from "@octokit/webhooks";

const cjkRegex = /[\u4E00-\u9FFF]/g;

export class ChineseTextSpamScanner implements SpamScanner {
  private readonly factor: number;

  constructor(factor: number) {
    this.factor = factor;
  }

  scanIssue(issue: WebhookPayloadIssuesIssue): boolean {
    return this.scanText(issue.body) || this.scanText(issue.title);
  }

  scanText(text: string): boolean {
    const matches = text.match(cjkRegex);
    return !!(matches != null && matches.length >= text.length * this.factor);
  }
}
