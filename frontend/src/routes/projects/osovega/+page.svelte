<script lang="ts">
    import type {IProjectItem} from '../../../interfaces/all';
    import ProjectHero from '../../../components/projects/ProjectHero.svelte';

    import {Projects} from '../../../stores/projects';
    import {page} from '$app/stores';

    const tokens = $page.url.pathname.split('/');
    const thisProject: IProjectItem = $Projects.find((p) => p.slug.includes(tokens[tokens.length - 1]))

</script>

<svelte:head>
    <title>Spence Wetjen</title>
    <meta name="description" content="Spence Wetjen | Full Stack Developer"/>
</svelte:head>

<section class="my-10">
    <ProjectHero project={thisProject}/>
</section>


<section>
    <div class="my-10 max-w-2xl mx-auto">
        <div class="mx-5 space-y-8">
            <h2>Key Accomplishments</h2>
            <div class="px-3">
                <ul class="list-disc list-outside">
                    <li>
                        Designed and developed a hybrid trading platform and seller marketplace that supported futures
                        contracts,
                        derivatives, and more.
                    </li>
                    <li> Developed a crop production tool for growers that integrated with the marketplace, helping
                        speculators predict future supply.
                    </li>
                    <li>
                        Led a team of four engineers to deliver a minimum viable product (MVP) in just ten months.
                    </li>
                </ul>
            </div>
        </div>
    </div>
    <div class="my-10 max-w-2xl mx-auto">
        <div class="mx-5 space-y-8">
            <h2>What was OsoVega?</h2>
            <p>
                OsoVega was an online marketplace where specialty growers and consumer-packaged-goods companies
                could buy and sell specialty agricultural crops. The platform prioritized helping growers to grade their
                crops and obtain the best possible prices, while also enabling buyers to purchase sufficient quantities
                to fulfill demand profitably.
            </p>
        </div>
    </div>

    <div class="my-10 max-w-2xl mx-auto">
        <div class="mx-5 space-y-8">
            <h2>Bridging the Gap Between Growers and Buyers</h2>
            <p>
                Our solution was a multi-sided marketplace. For the buyer side of the marketplace, a key problem is the
                need for a consistent supply at a predictable price. This is critical to maintaining margins and
                satisfying contracts with big retailers like Whole Foods.</p>
            <p>
                For the growers, who were the sellers on the marketplace, two analogies became apparent. First, growers
                are conditioned to lock in a price for their crop based on the movement of futures prices. Second, when
                growing a new crop, growers may require support with crop-specific farming practices, farm inputs, and
                harvesting.
            </p>
        </div>
    </div>

    <div class="my-10 max-w-2xl mx-auto">
        <div class="mx-5 space-y-8">
            <h2>Growers Dashboard</h2>
            <p>
                We opted for a hybrid e-commerce style interface that took inspiration from Amazon's seller site.
                Growers could list their products and set a fixed price or set a price range that they were willing to
                consider. Although we were not authorized to act as a Market Maker in the regulatory sense, we hired
                agents to connect with buyers and assist in facilitating the transaction all the way to delivery.
            </p>
            <h2>Predicting Crop Production</h2>
            <p>
                In low volume markets, the spot price of goods can vary wildly, making it difficult to ensure a
                consistent supply. To address this challenge, we developed a farm production tool that helps growers -
                and the marketplace as a whole - predict the volume of crops likely to be available in a given growing
                season.
            </p>
            <p>
                The production app tracked inputs and estimated yields based on feedback from specialists who visited
                the farms. This information helped growers plan their harvest and production schedules, as well as
                provide buyers with a more reliable estimate of future supply.
            </p>
        </div>
    </div>

    <div class="my-10 max-w-2xl mx-auto">
        <div class="mx-5 space-y-8">
            <h2>Technology</h2>
            <p>
                We chose to use C# DotNet and Postgres for the backend of the trading marketplace. The DotNet ecosystem
                for integrations and the compiler guarantees of C# made sense in the context of this product. The
                frontend was built in Vue and connected through Restful APIs.
            </p>
            <p>
                For the grower app, we decided to use Python with FastAPI. This application was primarily focused on
                CRUD operations with integrations to weather stations, drone data, and other sources. Python's
                flexibility and speed allowed us to deliver this part of the application quickly.
            </p>
            <p>
                We chose a monolithic frontend application all written in Vue. This approach provided a seamless user experience for growers and allowed us to easily integrate marketplace data with upcoming production data from the grower app.
            </p>
        </div>
    </div>

</section>
