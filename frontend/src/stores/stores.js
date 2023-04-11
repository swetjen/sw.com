import { writable } from 'svelte/store';

// ContactModal wraps the ContactModal state
function ContactModal() {
  const { subscribe, set, update } = writable(false);

  return {
    subscribe,
    toggle: () => update(n => !n),
    reset: () => set(false)
  };
}

const Contact = ContactModal();


// Contact Id Response
let ContactId = writable('')

if (typeof localStorage !== 'undefined') {
  // Get the value out of storage on load.
  const stored = localStorage.id
// or localStorage.getItem('content')

// Set the stored value or a sane default.
  ContactId = writable(stored || '')

// Anytime the store changes, update the local storage value.
  ContactId.subscribe((value) => localStorage.id = value)
// or localStorage.setItem('content', value)
}

export {ContactId, Contact}
