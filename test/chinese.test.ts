import { ChineseTextSpamScanner } from "../src/spam-scanner/chinese";

const scanner = new ChineseTextSpamScanner(0.4);

test("should not ban these texts", () => {
  expect(scanner.scanText("abc")).toBeFalsy();
  expect(scanner.scanText("abcggwp")).toBeFalsy();
  expect(scanner.scanText("好abcd")).toBeFalsy();
});

test("should ban these texts", () => {
  expect(scanner.scanText("好好好")).toBeTruthy();
  expect(scanner.scanText("好")).toBeTruthy();
  expect(scanner.scanText("好好好abcd")).toBeTruthy();
});
