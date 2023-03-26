<script>
  import { page } from '$app/stores';

  let crumbs

  // overrides the naive breadcrumbs implementation
  const overrides = [
    {
      "match": "projects",
      "href_instead": "/#projects"
    }
  ]


  $: {
    // Remove zero-length tokens.
    const tokens = $page.url.pathname.split('/').filter((t) => t !== '');

    // Create { label, href } pairs for each token.
    let tokenPath = '';
    crumbs = tokens.map((t) => {
      let override = overrides.find((e) => e.match = t);
      if (override) {
        return {
          label: t,
          href: override.href_instead
        }
      }
      tokenPath += '/' + t;
      return {
        label: t,
        href: tokenPath,
      };
    });

    // Add a way to get home too.
    crumbs.unshift({ label: 'home', href: '/' });
  }
</script>

{#each crumbs as c, i}
    {#if i == crumbs.length-1}
        <span class="capitalize">{c.label}</span>
    {:else}
        <a class="text-gray-500 cap" href={c.href}>{c.label}</a> <span class="text-gray-500">&gt;&nbsp;</span>
    {/if}
{/each}