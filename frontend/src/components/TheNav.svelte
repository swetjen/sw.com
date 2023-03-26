<script>
  import {page} from '$app/stores';
</script>

<header>
    <nav>
        <div>
            <div class="mx-auto flex justify-between items-center space-x-6 mt-4">
                <a class="flex-1 text-4xl font-bold" href="/">
                    Spence Wetjen
                </a>
                <a aria-current={$page.url.pathname === '/' ? 'page' : undefined} href="/">Home</a>
                <a aria-current={$page.url.pathname === '/' ? 'page' : undefined} href="/about">About</a>
                <a href="/">Contact</a>
            </div>
        </div>
    </nav>

</header>

