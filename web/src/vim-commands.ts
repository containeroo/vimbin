import { Vim } from "@replit/codemirror-vim";

Vim.defineEx("x", "x", function () {
  console.log("write");
});

Vim.defineEx("write", "w", function () {
  console.log("write");
});
