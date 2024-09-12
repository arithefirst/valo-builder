<script>
    import Skin from './Skin.svelte'
    export let weapon;
    let posts;

    async function fetchData() {
        const response = await fetch(`http://127.0.0.1:8080/api/v1/skin/${weapon}`);
        const data = await response.json();
        console.log(data);
        return data;
    }

    console.log(posts);

</script>

<main>
    {#await fetchData() then data}
        {#each data as skin}
            <div>
                <Skin src={skin.fullRender} title={skin.displayName}></Skin>
            </div>
        {/each}
    {/await}
</main>
