<script>
    export let src;
    export let title;
</script>

<div>
    <p><strong>{title}</strong></p>
    <img {src} alt="{title}">
</div>

<style>

    p {
        color: #cdd6f4;
    }

    div {
        border: 3px solid #6c7086;
        width: 250px;
        background-color: #1e1e2e;
        text-align: center;
        margin: 10px 5px;
        font-size: 0.8rem;
        border-radius: 15px;
        display: inline-grid;
    }

    img {
        width: 90%;
        padding: 10px;
    }

</style>