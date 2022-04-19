const { test, expect } = require("@playwright/test");
const path = require("path");

<<<<<<< HEAD
<<<<<<< HEAD
test("basic test", async ({ page }) => {
=======
test.beforeEach(async ({ page }) => {
>>>>>>> a4636229be3b4b37edbce94179d899e01a770c2c
=======
test.beforeEach(async ({ page }) => {
>>>>>>> 07b990f807137670d6b56e66abb172c46ab52015
  await page.goto(
    `file:${path.join(
      __dirname,
      "../../../../frontend/dom/guess-number-assignment/index.html"
    )}`
  );
<<<<<<< HEAD
<<<<<<< HEAD

  const title = await page.locator(".input").type("1");

  await page.locator(".check").click();

  if (
    (await page.locator(".input").inputValue()) <
    (await page.locator("#hidden-number").textContent())
  ) {
=======
=======
>>>>>>> 07b990f807137670d6b56e66abb172c46ab52015
});

test("check if input value smaller than random number", async ({ page }) => {
  await page.locator(".input").type("1");

  await page.locator(".check").click();
  const hidden_number = Number(
    await page.locator("#hidden-number").textContent()
  );
  expect(hidden_number).toBeGreaterThan(0);

  if (1 < hidden_number) {
<<<<<<< HEAD
>>>>>>> a4636229be3b4b37edbce94179d899e01a770c2c
=======
>>>>>>> 07b990f807137670d6b56e66abb172c46ab52015
    expect(await page.locator(".message").textContent()).toBe(
      "Too small, buddy!"
    );
  }
});
<<<<<<< HEAD
<<<<<<< HEAD
=======
=======
>>>>>>> 07b990f807137670d6b56e66abb172c46ab52015

test("check if input value greater than random number", async ({ page }) => {
  await page.locator(".input").type("10");

  await page.locator(".check").click();

  if (10 > (await page.locator("#hidden-number").textContent())) {
    expect(await page.locator(".message").textContent()).toBe(
      "Oww... that's too big!"
    );
  }
});

test("check if input value smaller than 1", async ({ page }) => {
  await page.locator(".input").type("0");

  await page.locator(".check").click();

  expect(await page.locator(".message").textContent()).toBe(
    "Guess any number between 1 and 10"
  );
});

test("check if input value greater than 10", async ({ page }) => {
  await page.locator(".input").type("11");

  await page.locator(".check").click();

  expect(await page.locator(".message").textContent()).toBe(
    "Guess any number between 1 and 10"
  );
});

test("reset message", async ({ page }) => {
  await page.locator(".new-game").click();

  expect(await page.locator(".message").textContent()).toBe(
    "Lets start guessing..."
  );
});
<<<<<<< HEAD
>>>>>>> a4636229be3b4b37edbce94179d899e01a770c2c
=======
>>>>>>> 07b990f807137670d6b56e66abb172c46ab52015
