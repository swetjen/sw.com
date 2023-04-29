<script lang="ts">
    import type {IProjectItem} from '../../../interfaces/all';
    import ProjectHero from '../../../components/projects/ProjectHero.svelte';

    import {Projects} from '../../../stores/projects';
    import {page} from '$app/stores';
    import TechLogo from "../../../components/TechLogo.svelte";
    import Image from "../../../components/Image.svelte";
    import Caption from "../../../components/Caption.svelte";

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
                        contracts, and product delivery.
                    </li>
                    <li> Developed a crop production tool for growers that integrated with the marketplace, helping
                        market participants estimate future supply.
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
                growing a new crop, growers may require support with farming practices, farm inputs, and
                harvesting.
            </p>
        </div>
    </div>
</section>

<section>
    <div class="flex flex-row space-x-4 justify-around bg-gray-300 rounded p-4 my-2 w-full">
        <div class="space-y-4">
            <Image class="max-w-sm" src="/ov/Edit Season.jpg"
                   alt="Wireframe of creating a season in the Osovega Growers application"/>
        </div>
        <div class="space-y-4">
            <Image class="max-w-md" src="/ov/Growers Dashboard.png"
                   alt="Wireframe of creating a season in the Osovega Growers application"/>
            <Image class="max-w-md" src="/ov/Growers Inventory.png"
                   alt="Wireframe of creating a season in the Osovega Growers application"/>
        </div>
        <!--        <div>-->
        <!--            <Image class="max-w-md" src="/ov/Edit Harvest.jpg"-->
        <!--                   alt="Wireframe of creating a season in the Osovega Growers application"/>-->
        <!--        </div>-->
    </div>
    <div>
        <Caption
                props="Osovega — Wireframes illustrating harvest planning and the grower's sell-side interface."/>
    </div>
</section>

<section>
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
                consistent supply. To address this challenge, we developed a farm production tool that helped growers -
                and the marketplace as a whole - predict the volume of crops likely to be available in a given growing
                season.
            </p>
        </div>
    </div>
</section>

<section>
    <div class="flex flex-row space-x-4 justify-around bg-gray-300 rounded p-4 my-2 w-full">
        <div class="space-y-4">
            <Image class="max-w-sm" fullSrc="/ov/Dashboard.jpg" src="/ov/GrowersApp.jpg"
                   alt="Wireframe of creating a season in the Osovega Growers application"/>
        </div>
    </div>
    <div>
        <Caption
                props="Osovega — Wireframe illustrating farm production tool with harvest prediction."/>
    </div>
</section>
<section>
    <div class="my-10 max-w-2xl mx-auto">
        <div class="mx-5 space-y-8">
            <p>
                The production app tracked inputs and estimated yields based on feedback from specialists who
                visited
                the farms. This information helped growers plan their harvest and production schedules, as well as
                provide buyers with a more reliable estimate of future supply.
            </p>
        </div>
    </div>

    <div class="my-10 max-w-2xl mx-auto">
        <div class="mx-5 space-y-8">
            <h2>Technology</h2>
            <div class="grid fill-current text-gray-700 gap-2 grid-flow-col grid-cols-3 md:grid-cols-6 grid-rows-2 md:grid-rows-1 items-center">
                <div class="text-sm p-2 bg-gray-50 rounded h-18 space-x-4 flex flex-col h-full justify-between items-center">
                    <div class="w-8 max-w-full my-auto text-sm mx-auto">
                        <TechLogo name="C#"/>
                    </div>
                </div>
                <div class="text-sm p-2 bg-gray-50 rounded h-18 space-x-2 flex flex-col h-full justify-between items-center">
                    <div class="w-12 max-w-full my-auto">
                        <TechLogo class="w-8 mx-auto my-auto" name="DotNet"/>
                    </div>
                </div>
                <div class="text-sm p-2 bg-gray-50 rounded h-18 space-x-2 flex h-full flex-col justify-between items-center">
                    <div class="w-6 max-w-full">
                        <TechLogo name="Python 3"/>
                    </div>
                    Python
                </div>
                <div class="text-sm p-2 bg-gray-50 rounded h-18 space-x-2 flex h-full flex-col justify-between items-center">
                    <div class="w-6 max-w-full">
                        <TechLogo name="Postgres"/>
                    </div>
                    Postgres
                </div>
                <div class="text-sm p-2 bg-gray-50 rounded h-18 space-x-2 flex flex-col h-full justify-between items-center">
                    <div class="w-8 max-w-full">
                        <TechLogo name="Typescript"/>
                    </div>
                    Typescript
                </div>
                <div class="text-sm p-2 bg-gray-50 rounded h-18 space-x-2 flex flex-col h-full justify-between items-center">
                    <div class="w-8 max-w-full">
                        <TechLogo name="Vue"/>
                    </div>
                    Vue
                </div>
            </div>
            <p>
                We chose to use C# DotNet and Postgres for the backend of the trading marketplace. The DotNet
                ecosystem
                for integrations and the language level compiler guarantees made sense in the context of this
                product. A
                series of Restful APIs exposed the backend service.
            </p>
            <p>
                For the farm production app, we decided to use Python with FastAPI for the backend. This application
                was
                primarily focused on CRUD operations with integrations to weather stations, drone data, and other
                sources. Python's flexibility and speed allowed us to deliver this part of the application quickly.
            </p>
            <p>
                We chose a monolithic frontend application all written in Vue. This approach provided a seamless
                user
                experience for growers and allowed us to easily integrate marketplace data with production data
                coming
                from the fields.
            </p>
        </div>
    </div>

</section>
