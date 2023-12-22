document.addEventListener("DOMContentLoaded", function () {
  // Function to get the preferred theme
  function getPreferredTheme() {
    const prefersDarkMode = window.matchMedia?.(
      "(prefers-color-scheme: dark)",
    )?.matches;

    // If theme is set to "auto" and the system prefers dark mode, return "darkTheme"
    if (theme === "auto" && prefersDarkMode) {
      return darkTheme;
    }

    const prefersLightMode = window.matchMedia?.(
      "(prefers-color-scheme: light)",
    )?.matches;

    // If theme is set to "auto" and the system prefers dark mode, return "darkTheme"
    if (theme === "auto" && prefersLightMode) {
      return lightTheme;
    }

    console.log(`Theme set to '${theme}'`);
    // For any other case, return the actual theme
    return theme;
  }

  // Function to set the theme based on the initial color scheme or the 'theme' variable
  function setThemeBasedOnColorScheme() {
    const preferredTheme = getPreferredTheme();
    console.log(`Setting theme to '${preferredTheme}'`);
    editor.setOption("theme", preferredTheme);
  }
  // Function to update Vim mode display
  function updateVimMode(vimEvent, vimModeElement) {
    const { mode, subMode } = vimEvent;

    console.log(`VIM mode '${mode}', subMode '${subMode}'`);

    // Mapping of mode to corresponding text and class
    const modeMap = {
      normal: { text: "NORMAL", class: "normal" },
      insert: { text: "INSERT", class: "insert" },
      visual: {
        text: subMode === "" ? "VISUAL" : "V-LINE",
        class: subMode === "" ? "visual" : "visual-line",
      },
      unknown: { text: "UNKNOWN", class: "unknown" },
    };

    // Remove all existing classes
    vimModeElement.classList.remove(
      ...Object.values(modeMap).map((entry) => entry.class),
    );

    // Update text and add corresponding class
    const { text, class: modeClass } = modeMap[mode] || modeMap.unknown;
    vimModeElement.innerText = text;
    vimModeElement.classList.add(modeClass);
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
    statusElement.classList.remove("isError", "noChanges");

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
    lineWrapping: true,
  });

  editor.on("cursorActivity", showRelativeLines);

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
  window
    .matchMedia("(prefers-color-scheme: dark)")
    .addListener(setThemeBasedOnColorScheme);
  window
    .matchMedia("(prefers-color-scheme: light)")
    .addListener(setThemeBasedOnColorScheme);

  // Focus editor
  editor.focus();
});
