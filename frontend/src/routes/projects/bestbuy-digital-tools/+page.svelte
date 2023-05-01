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


<section class="my-10 max-w-2xl mx-auto">
    <div class="mx-5 space-y-8">
        <h2>
            Key Achievements
        </h2>
        <div class="px-3">
            <ul class="list-disc list-outside">
                <li>
                    Led the transformation of the company's digital tools by breaking down silos, creating transparency,
                    and aligning teams and budgets.
                </li>
                <li>Successfully managed a two-year, tens-of-millions dollar project to reorganize and align the digital
                    tools, resulting in significant cost savings.
                </li>
                <li>Achieved a favorable ongoing return on investment (ROI) by reducing operating expenses and outside
                    spend.
                </li>
            </ul>
        </div>

        <p>Best Buy, a major retailer, was expending a significant amount of money on redundant apps, websites, and
            other digital tools, without any coherent plan or coordination among teams. This caused confusion for both
            customers and employees, resulting in significant costs for the company.
        </p>
        <p>
            The Digital Tools project addressed this issue by consolidating all the work under a unified structure and
            utilizing enterprise capabilities to optimize the value and impact of these tools. As a result, the company
            was able to save a considerable amount of money, enable employees to perform at their best, and enhance the
            shopping experience for customers at Best Buy.
        </p>
        <h2>Changing Status Quo</h2>
        <p>
            It all started when I discovered yet another product finder. Product finders were popular in the late
            2000's, but this one was terrible. It broke easily, recommended the wrong products and didn't even know what
            we had in inventory. Worst, the company had overpaid for it.
        </p>
        <p>
            Fed up with the status quo I set to changing the status quo. I began by mapping out all of Best Buy's
            digital tools, focusing on overlapping concern, gap areas and cost. with emphasis to overlapping concerns
            and cost. Then came the hard part of getting people on board.
        </p>
        <p>
            Delivering a project of this magnitude in a large corporation requires both diplomacy and tenacity. The
            delicate part is helping teams transition from a state of caution to being supportive by demonstrating how
            their work and objectives align with the outcome and business benefits being pursued. At the same time,
            tenacity is needed as the teams that need to support the project have competing priorities and may be
            stretched thin.
        </p>
        <h2>Lessons Learned Getting People On-board</h2>
        <div class="px-3">
            <ul class="list-disc list-outside">
                <li>
                    Clearly define the problem at hand
                </li>
                <li>
                    Gain agreement that the problem affects everyone
                </li>
                <li>Reach an agreement that the benefits of solving the problem outweigh the potential risks and other
                    competing priorities.
                </li>
            </ul>
        </div>
            <div style="padding:75% 0 0 0;position:relative;"><iframe src="https://player.vimeo.com/video/822532855?h=f2bcbb9aa2&amp;badge=0&amp;autopause=0&amp;player_id=0&amp;app_id=58479" frameborder="0" allow="autoplay; fullscreen; picture-in-picture" allowfullscreen style="position:absolute;top:0;left:0;width:100%;height:100%;" title="digtal-tools"></iframe></div><script src="https://player.vimeo.com/api/player.js"></script>
        <div class="my-10 max-w-2xl mx-auto">
            <div class="mx-5 space-y-8">
                <h2>Check it out</h2>
                <p>Visit <a target="_blank" href="https://bestbuy.com">BestBuy.com</a> online or a store near you.</p>
            </div>
        </div>
    </div>
</section>
