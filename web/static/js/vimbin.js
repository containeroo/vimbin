// Desc: Main JavaScript file for the Vimbin web app
document.addEventListener("DOMContentLoaded", function () {
  // Function to get the preferred theme (dark, light, or system default)
  function getPreferredTheme() {
    // If the browser doesn't support the prefers-color-scheme media query, return the default theme
    if (!window.matchMedia) {
      return "default";
    }

    if (window.matchMedia("(prefers-color-scheme: dark)").matches) {
      return "catppuccin";
    }

    if (window.matchMedia("(prefers-color-scheme: light)").matches) {
      return "default";
    }

    return "default";
  }

  // Function to update Vim mode display
  function updateVimMode(vimEvent, vimModeElement) {
    const mode = vimEvent.mode;
    const sub = vimEvent.subMode;

    vimModeElement.classList.remove(
      "normal",
      "insert",
      "visual",
      "visual-line",
    );

    switch (mode) {
      case "normal":
        vimModeElement.innerText = "NORMAL";
        vimModeElement.classList.add("normal");
        break;
      case "insert":
        vimModeElement.innerText = "INSERT";
        vimModeElement.classList.add("insert");
        break;
      case "visual":
        if (sub === "") {
          vimModeElement.innerText = "VISUAL";
          vimModeElement.classList.add("visual");
          break;
        }
        vimModeElement.innerText = "V-LINE";
        vimModeElement.classList.add("visual-line");
        break;
      default:
        vimModeElement.innerText = "UNKNOWN";
        vimModeElement.classList.add("unknown");
    }
  }

  // Function to show relative line numbers
  function showRelativeLines(cm) {
    const lineNum = cm.getCursor().line + 1; // Get the current line number of the cursor

    // If the current line number is the same as the stored line number, no need to update
    if (cm.state.curLineNum === lineNum) {
      return;
    }

    cm.state.curLineNum = lineNum; // Update the stored line number

    // Set the line number formatter to display relative line numbers
    cm.setOption("lineNumberFormatter", (l) =>
      // If the line number is the same as the current line, display the absolute line number
      l === lineNum ? lineNum : Math.abs(lineNum - l),
    );
  }

  // Function to save the content
  async function saveContent() {
    const statusElement = document.getElementById("status");
    clearTimeout(statusElement.timerId); // Clear the existing timer before setting a new one

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
        body: JSON.stringify({ content: editor.getValue() }),
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
      status = `ERROR: ${error.message}`;
    }

    statusElement.innerText = status;
    statusElement.classList.remove("isError", "noChanges"); // Remove all classes

    if (isError) {
      statusElement.classList.add("isError");
    }

    if (noChanges) {
      statusElement.classList.add("noChanges");
    }

    if (startTimer) {
      const delay = 5000;

      // Set a new timer
      statusElement.timerId = setTimeout(() => {
        statusElement.innerText = "";
        statusElement.classList.remove("isError", "noChanges");
      }, delay);
    }
  }

  var editor = CodeMirror.fromTextArea(document.getElementById("code"), {
    lineNumbers: true,
    mode: "text/x-csrc",
    keyMap: "vim",
    matchBrackets: true,
    showCursorWhenSelecting: true,
    theme: getPreferredTheme(),
    lineWrapping: true, // Optional: enable line wrapping if desired
  });

  editor.on("cursorActivity", showRelativeLines);
  editor.focus();

  // Custom vim Ex commands
  CodeMirror.Vim.defineEx("x", "", function () {
    saveContent();
  });

  var vimMode = document.getElementById("vim-mode");
  CodeMirror.on(editor, "vim-mode-change", function (e) {
    updateVimMode(e, vimMode);
  });

  CodeMirror.commands.save = saveContent;

  // Listen for changes in the prefers-color-scheme media query
  window.matchMedia("(prefers-color-scheme: dark)").addListener((e) => {
    if (e.matches) {
      editor.setOption("theme", "catppuccin");
    } else {
      editor.setOption("theme", "default");
    }
  });

  window.matchMedia("(prefers-color-scheme: light)").addListener((e) => {
    if (e.matches) {
      editor.setOption("theme", "default");
    }
  });
});
