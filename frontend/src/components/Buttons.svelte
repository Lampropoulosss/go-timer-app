<script>
  import { createEventDispatcher } from "svelte";
  import { scale } from "svelte/transition";

  const dispatch = createEventDispatcher();

  export let timerStarted = false;
  export let isPaused = false;
  export let alarmFinished = false;

  export let isInvalid = false;
  export let timer = "";

  function validateInput() {
    // Regular expression to validate HH:MM:SS format
    const regex = /^([0-1]\d|2[0-3]):([0-5]\d):([0-5]\d)$/;
    isInvalid = !regex.test(timer);
  }
</script>

{#if timerStarted && !alarmFinished}
  <button
    in:scale={{ duration: 250 }}
    on:click={() => {
      dispatch("pause-resume");
      isPaused = !isPaused;
    }}>{isPaused ? "Resume" : "Pause"}</button
  >
  <button
    in:scale={{ duration: 250 }}
    on:click={() => {
      isInvalid = false;
      dispatch("clear");
    }}>Stop</button
  >
{:else if timerStarted && alarmFinished}
  <button
    in:scale={{ duration: 250 }}
    on:click={() => {
      isInvalid = false;
      dispatch("clear");
    }}>Stop</button
  >
{:else}
  <button
    in:scale={{ duration: 250 }}
    on:click={() => {
      validateInput();
      if (isInvalid) {
        return dispatch("invalid");
      }
      dispatch("start", { timer });
    }}>Start</button
  >
{/if}

<style>
  button {
    border-radius: 8px;
    border: 2px solid transparent;
    padding: 0.4em 0.5em;
    font-size: 1em;
    font-weight: 500;
    font-family: inherit;
    color: #ffffff;
    background-color: #0f0f0f98;
    transition: border-color 0.25s;
    box-shadow: 0 2px 2px rgba(0, 0, 0, 0.2);
    cursor: pointer;
    outline: none;
  }

  button:hover {
    border-color: #00ced1;
    background-color: #0f0f0f59;
  }
</style>
