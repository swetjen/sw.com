<script lang="ts">
  import type {IProjectItem} from '../../interfaces/all';
  import { createEventDispatcher } from 'svelte';

  const dispatch = createEventDispatcher()

  function sendCloseMsg() {
    console.log('send called')
    dispatch('message', {
      text: 'click'
    })
  }

  export let props: IProjectItem

</script>

<div class="group relative rounded-lg text-sm leading-6 wiggle">
    <div class="h-24 bg-center bg-cover rounded-lg group-hover:bg-white" style={`background-image: url('${props.heroImage}')`}>
        <div class="flex flex-col h-full justify-center items-center bg-center bg-cover rounded-2xl transform-gpu hover:scale-95 duration-150 ease-in-out shadow-2xl hover:shadow-3xl ">
            <img class="inline-block mx-auto opacity-90" width="48" src={props.logo}>
        </div>
    </div>
    <a on:click={sendCloseMsg} href={props.slug} class="block text-gray-900">
        {props.title}
        <span class="absolute inset-0"></span>
    </a>
</div>


<style>

    @keyframes wobble {
        0% { transform: translateX(0%); }
        15% { transform: rotate(-4deg); }
        30% { transform: rotate(2deg); }
        45% { transform:  rotate(-2deg); }
        60% { transform: rotate(1deg); }
        75% { transform:  rotate(-1deg); }
        100% { transform: translateX(0%); }
    }

    .wiggle:hover {
        animation-name: wobble;
        animation-duration: 800ms;
        animation-timing-function: linear;
    }

</style>