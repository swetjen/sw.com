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

export const Contact = ContactModal();
