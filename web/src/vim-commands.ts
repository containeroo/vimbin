import { Vim } from "@replit/codemirror-vim";
import { editor } from "./editor-setup";
import { apiToken } from "./main";

// Function to save the content
async function saveContent() {
  const content = editor.state.doc.toString();

  const statusElement = document.getElementById("status") as HTMLElement & {
    timerId?: number;
  };

  // check if statusElement is null
  if (!statusElement) {
    return;
  }

  clearTimeout(statusElement.timerId);

  let status = "No changes were made.";
  let isError = false;
  let noChanges = true;
  let startTimer = true;

  try {
    const response = await fetch("/save", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        "X-API-Token": apiToken,
      },
      body: JSON.stringify({ content: content }),
    });

    if (!response.ok) {
      throw new Error(`Save failed. Reason: ${response.statusText}`);
    }

    const isJson = response.headers
      .get("content-type")
      ?.includes("application/json");

    if (!isJson) {
      throw new Error("Response was not JSON");
    }

    const changesResponse = await response.json();

    if (changesResponse.status !== "no changes") {
      const bytesWritten = response.headers.get("X-Bytes-Written");
      status = `${bytesWritten}B written`;
      noChanges = false;
    }
  } catch (error) {
    startTimer = false;
    isError = true;
    status = `ERROR: ${(error as Error).message}`;
  }

  statusElement.innerText = status;
  statusElement.classList.remove("isError", "noChanges");

  if (isError) {
    statusElement.classList.add("isError");
  }

  if (noChanges) {
    statusElement.classList.add("noChanges");
  }

  if (startTimer) {
    const delay = 5000;

    statusElement.timerId = setTimeout(() => {
      statusElement.innerText = "";
      statusElement.classList.remove("isError", "noChanges");
    }, delay);
  }
}

// Vim.unmap("<space>");
// Vim.map("<space>", "<leader>");

Vim.defineEx("x", "x", function () {
  saveContent();
});

Vim.defineEx("write", "w", function () {
  saveContent();
});

Vim.map("<Space><Space>", "l"); // Move right

// Define a custom action for yanking selected text
Vim.defineAction("yankToClipboard", function () {
  const text = editor.state.sliceDoc(
    editor.state.selection.main.from,
    editor.state.selection.main.to,
  );

  let status = "";
  let isError = false;
  let startTimer = true;

  const statusElement = document.getElementById("status") as HTMLElement & {
    timerId?: number;
  };

  // check if statusElement is null
  if (!statusElement) {
    return;
  }

  clearTimeout(statusElement.timerId);

  navigator.clipboard
    .writeText(text)
    .then(() => {
      status = "Text yanked to clipboard";

      console.log(status);
    })
    .catch((err) => {
      isError = true;
      startTimer = false;
      status = `ERROR: ${(err as Error).message}`;

      console.error(status);
    });

  statusElement.innerText = status;

  if (isError) {
    statusElement.classList.add("isError");
  } else {
    statusElement.classList.remove("isError");
  }

  if (startTimer) {
    const delay = 5000;

    statusElement.timerId = setTimeout(() => {
      statusElement.innerText = "";
      statusElement.classList.remove("isError");
    }, delay);
  }
});
Vim.mapCommand("<Space>y", "action", "yankToClipboard"); // Map the custom action to a command triggered by `<Space>y`

// // Define a custom action for yanking the current line
// Vim.defineAction("yankCurrentLine", function () {
//   const lineNumber = editor.state.selection.main.anchor; // Get the line number where the cursor currently is
//   const lineContent = editor.state.doc.lineAt(lineNumber).text;
//
//   // Use the Clipboard API to copy the line's content
//   navigator.clipboard
//     .writeText(lineContent)
//     .then(() => {
//       console.log("Current line yanked to clipboard");
//     })
//     .catch((err) => {
//       console.error("Failed to copy current line: ", err);
//     });
// });
// Vim.mapCommand("<Space>yy", "action", "yankCurrentLine"); // Map the custom action to a command triggered by `<Space>yy`

// Optionally, if you want to redefine the leader key or ensure it's set to space
Vim.unmap("<Space>"); // Unmap existing mappings for space if necessary
// Here you could define `<Space>` as a prefix or leader for further commands as needed
