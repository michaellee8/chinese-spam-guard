import { Application } from "probot/lib/application";
import { ChineseTextSpamScanner } from "./spam-scanner/chinese"; // eslint-disable-line no-unused-vars

const scanner = new ChineseTextSpamScanner(0.4);

export = (app: Application) => {
  app.on("issues.opened", async context => {
    if (scanner.scanIssue(context.payload.issue)) {
      await context.github.issues.update(context.issue({ state: "closed" }));
    }
  });
};
