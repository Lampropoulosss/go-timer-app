<script>
  // @ts-nocheck

  import { Notify, Stop } from "../wailsjs/go/main/App.js";
  import Options from "./components/Options.svelte";
  import { writable } from "svelte/store";
  import { onMount } from "svelte";
  import { shortcutDisabler } from "./utilities/shortcutDisabler";

  shortcutDisabler(); // Disable shortcuts such as ctrl + R etc.

  let timer = writable(0);
  let interval;
  // let audio;

  let isPaused = false;
  let alarmFinished = false;

  function formatTime(totalSeconds) {
    const hours = String(Math.floor(totalSeconds / 3600)).padStart(2, "0");
    const minutes = String(Math.floor((totalSeconds % 3600) / 60)).padStart(
      2,
      "0"
    );
    const seconds = String(totalSeconds % 60).padStart(2, "0");
    return `${hours}:${minutes}:${seconds}`;
  }

  // @ts-ignore
  function startTimer(event) {
    const [hours, minutes, seconds] = event.detail.timer.split(":").map(Number);
    let totalSeconds = hours * 3600 + minutes * 60 + seconds;

    timer.set(totalSeconds);

    if (interval) {
      clearInterval(interval);
    }

    interval = setInterval(() => {
      if (isPaused) return;

      totalSeconds -= 1;
      timer.set(totalSeconds);

      if (totalSeconds <= 0) {
        clearInterval(interval);
        alarmFinished = true;
        timer.set(0);
        Notify();
      }
    }, 1000);

    return () => clearInterval(interval);
  }

  function clearTimer() {
    clearInterval(interval);
    timer.set(0);
    alarmFinished = false;
    Stop();
  }

  function pause_resume() {
    isPaused = !isPaused;
  }
</script>

<div class="container">
  <div class="timer" class:colored={alarmFinished}>
    {#if $timer > 0}
      {formatTime($timer)}
    {:else}
      00:00:00
    {/if}
  </div>
  <Options
    on:start={startTimer}
    on:clear={clearTimer}
    on:pause-resume={pause_resume}
    {isPaused}
    {alarmFinished}
    timerStarted={$timer > 0 || alarmFinished}
  />
</div>

<style>
  @font-face {
    font-family: "Pacifico";
    font-style: normal;
    font-weight: 400;
    src: url("./assets/fonts/Pacifico-Regular.ttf") format("truetype");
  }

  :root {
    font-family: cursive, sans-serif;
    font-size: 16px;
    line-height: 24px;
    font-weight: 400;

    color: #f6f6f6;
    background-color: #2f2f2f;

    font-synthesis: none;
    text-rendering: optimizeLegibility;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    -webkit-text-size-adjust: 100%;

    -webkit-touch-callout: none; /* iOS Safari */
    -webkit-user-select: none; /* Safari */
    -khtml-user-select: none; /* Konqueror HTML */
    -moz-user-select: none; /* Old versions of Firefox */
    -ms-user-select: none; /* Internet Explorer/Edge */
    user-select: none; /* Non-prefixed version, currently supported by Chrome, Edge, Opera and Firefox */
  }

  :global(*) {
    box-sizing: border-box;
  }

  :global(body) {
    margin: 0;
    padding: 0;
  }

  .container {
    min-height: 100vh;
    display: flex;
    align-items: center;
    flex-direction: column;
    justify-content: center;
    text-align: center;
  }

  .colored {
    color: #00ced1;
  }

  .timer {
    font-family: Pacifico;
    font-size: 3em;
    font-weight: 700;
    margin-bottom: 0.4em;
  }
</style>
