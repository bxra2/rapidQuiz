<script lang="ts">
  import Button from "./Button.svelte";
  import { createEventDispatcher } from "svelte";
  const dispatch = createEventDispatcher();
  export let quiz: { _id: string; name: string };

  function host() {
    dispatch("host");
  }
</script>

<div
  class="flex justify-between items-center bg-white border-2 border-green-500 p-4 rounded-xl"
>
  {quiz.name}
  <Button on:click={host}>Host</Button>
</div>
