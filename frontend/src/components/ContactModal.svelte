<script lang="ts">
  import type {IContactForm, IContactFormResponse} from '../interfaces/all';
  import {Contact, ContactId} from '../stores/stores';
  import { toast } from '@zerodevx/svelte-toast';
  import { fly, fade } from 'svelte/transition';

  // state
  let errors: IContactForm = {
    name: '',
    message: '',
    email: '',
  }

  function isValid(): boolean {
    let valid = true
    if (!form.email) {
      valid = false
      errors.email = "An email address is required"
      return valid
    }
    errors.email = ""
    if (form.message.length < 3) {
      valid = false
      errors.message = "Please include a message."
      return valid
    }
    errors.message = ''
    return valid
  }

  let is_valid = true
  let is_submitting = false

  //form
  let form: IContactForm = {
    name: '',
    message: '',
    email: ''
  }

  let result = null

  async function SubmitContactForm() {
    if (!isValid()) {
      return
    }
    is_submitting = true
      // todo : handle environments
    const res = await fetch('https://spencewetjen.com/api/v1/contact', {
      method: 'POST',
      body: JSON.stringify(form
      )
    })
    if (res.status === 200) {
      const data = await res.json() as IContactFormResponse

      $ContactId = data.guid

      result = data.status
      toast.push("Success", {
        duration: 3500,
        theme: {
          '--toastColor': 'mintcream',
          '--toastBackground': 'rgba(72,187,120,0.9)',
          '--toastBarBackground': '#2F855A'
        }});
      $: Contact.toggle()
    }
    is_submitting = false
    // Handle errors
  }

</script>


{#if $Contact}
<div>
    <div class="fixed z-10 inset-0 overflow-y-auto">
            <div
                class="flex  md:items-end justify-center  pt-4 px-4 pb-20 text-center sm:block sm:p-0">

                <!--
              Background overlay, show/hide based on modal state.

              Entering: "ease-out duration-300"
                From: "opacity-0"
                To: "opacity-100"
              Leaving: "ease-in duration-200"
                From: "opacity-100"
                To: "opacity-0"
            -->
                <div in:fade="{{ duration: 300}}" out:fade="{{duration: 200}}" class="fixed inset-0 transition-opacity" aria-hidden="true">
                    <div
                            on:click={Contact.reset}
                            class="absolute inset-0 bg-gray-900 opacity-75"></div>
                </div>

                <!-- This element is to trick the browser into centering the modal contents. -->
                <span class="hidden sm:inline-block sm:align-middle sm:h-screen" aria-hidden="true">&#8203;</span>
                <!--
                  Modal panel, show/hide based on modal state.

                  Entering: "ease-out duration-300"
                    From: "opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
                    To: "opacity-100 translate-y-0 sm:scale-100"
                  Leaving: "ease-in duration-200"
                    From: "opacity-100 translate-y-0 sm:scale-100"
                    To: "opacity-0 translate-y-4 sm:translate-y-0 sm:scale-95"
                -->
                <div in:fly="{{ duration: 300, y: 18}}" out:fly="{{duration:200, y:18}}"
                        class="inline-block align-bottom bg-white rounded-lg text-left overflow-hidden shadow-xl transform transition-all sm:my-8 sm:align-middle sm:max-w-lg sm:w-full"
                        role="dialog" aria-modal="true" aria-labelledby="modal-headline">
                    <div class="bg-white px-4 pt-5 pb-4 sm:p-6 sm:pb-4">
                        <div class="sm:flex sm:items-start">
                            <!--            <div-->
                            <!--              class="mx-auto flex-shrink-0 flex items-center justify-center h-12 w-12 rounded-full bg-gray-100i sm:mx-0 sm:h-10 sm:w-10">-->
                            <!--              &lt;!&ndash; Heroicon name: outline/exclamation &ndash;&gt;-->
                            <!--              <svg class="h-6 w-6 text-black" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24"-->
                            <!--                   stroke="currentColor" aria-hidden="true">-->
                            <!--                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"-->
                            <!--                      d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"/>-->
                            <!--              </svg>-->
                            <!--            </div>-->
                            <div class="mt-3 text-center sm:mt-0 sm:ml-4 sm:text-left">
                                <h3 class="text-lg leading-6 font-medium text-gray-900" id="modal-headline">
                                    Contact Spence
                                </h3>
                                <div class="mt-2">
                                    <p class="text-sm text-gray-500">
                                        Interested to work together? Fill out the from below and I'll be in touch.
                                    </p>
                                </div>
                            </div>
                        </div>
                        <div>
                            <div class="px-4 pb-4">
                                <!-- Contact Form -->
                                <form class="mt-8 space-y-8">
                                    <div class="relative mt-2 rounded-md shadow-sm">
                                            <label for="email-address" class="sr-only">Email address</label>
                                            <input id="email-address" bind:value={form.email} name="email" type="email" autocomplete="email" required
                                                   placeholder="Email address"
                                                   disabled={is_submitting}
                                                   class="{errors.email ? 'ring-red-300 placeholder:text-red-300 focus:ring-red-500 text-red-900' : 'text-gray-900 ring-gray-300 placeholder:text-gray-400 focus:ring-indigo-600' } block w-full rounded-md border-0 py-1.5 pr-10 ring-1 ring-inset focus:ring-2 focus:ring-inset sm:text-sm sm:leading-6">
                                            <div class="pointer-events-none absolute inset-y-0 right-0 flex items-center pr-3">
                                                <svg class="h-5 w-5 text-red-500" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
                                                    <path fill-rule="evenodd" d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-8-5a.75.75 0 01.75.75v4.5a.75.75 0 01-1.5 0v-4.5A.75.75 0 0110 5zm0 10a1 1 0 100-2 1 1 0 000 2z" clip-rule="evenodd" />
                                                </svg>
                                            </div>
                                        {#if errors.email}
                                        <p class="mt-2 text-sm text-red-600" id="email-error">{errors.email}</p>
                                        {/if}
                                        <div>
                                            <label for="message" class="sr-only">Email address</label>
                                            <textarea disabled={is_submitting}
                                                      id="message" name="message" bind:value={form.message} type="text" required
                                                      placeholder="Message"
                                                      rows="8"
                                                      class="{errors.message ? 'ring-red-300 placeholder:text-red-300 focus:ring-red-500 text-red-900' : 'text-gray-900 ring-gray-300 placeholder:text-gray-400 focus:ring-indigo-600' }  w-full appearance-none rounded-none relative block px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-b-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10 sm:text-sm"
                                            ></textarea>
                                        </div>
                                        {#if errors.message}
                                            <p class="mt-2 text-sm text-red-600" id="message-error">{errors.message}</p>
                                        {/if}
                                    </div>
                                </form>
                            </div>
                        </div>
                    </div>
                    <div class="bg-gray-100 py-3 md:px-10 px-6 sm:flex justify-end flex-row ">

                        <button type="button"
                                on:click={Contact.reset}
                                class="mt-3 inline-flex justify-center rounded-md border border-gray-300 shadow-sm px-4 py-2 bg-white text-base font-medium text-gray-700 hover:bg-gray-50 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500 sm:mt-0 sm:ml-3 sm:w-auto sm:text-sm">
                            Cancel
                        </button>

                        <button type="button" on:click={SubmitContactForm}
                                class:bg-green-500={is_valid}
                                class="w-24 bg-gray-300 inline-flex justify-center rounded-md border border-transparent shadow-sm px-4 py-2 text-base font-medium text-white focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-red-500 sm:ml-3 sm:w-auto sm:text-sm">
                            {#if !is_submitting}
                            <span>Submit</span>
                            {:else}
                            <span class="w-full flex items-center">
                <div class="w-5 flex-none" id="gear">
                  <svg aria-hidden="true" focusable="false" data-prefix="fas" data-icon="spinner-third"
                       class="svg-inline--fa fa-spinner-third fa-w-16" role="img" xmlns="http://www.w3.org/2000/svg"
                       viewBox="0 0 512 512">
                    <path fill="currentColor"
                          d="M456.433 371.72l-27.79-16.045c-7.192-4.152-10.052-13.136-6.487-20.636 25.82-54.328 23.566-118.602-6.768-171.03-30.265-52.529-84.802-86.621-144.76-91.424C262.35 71.922 256 64.953 256 56.649V24.56c0-9.31 7.916-16.609 17.204-15.96 81.795 5.717 156.412 51.902 197.611 123.408 41.301 71.385 43.99 159.096 8.042 232.792-4.082 8.369-14.361 11.575-22.424 6.92z"></path>
                  </svg>
                </div>
              </span>
                            {/if}
                        </button>
                    </div>
                </div>

            </div>
    </div>
</div>
{/if}

<style>

    #gear {
        animation-name: sp;
        animation-duration: 0.75s;
        animation-iteration-count: infinite;
        transform-origin: 50% 50%;
    }

    @keyframes sp {
        0% {
            transform: rotate(0deg);
        }
        100% {
            transform: rotate(360deg);
        }
    }

</style>
