<template>
    <ul style="list-style: none; padding-left: 0;">
        <li>
            <img src="/phone.png" class="icon">
            <span :contenteditable="editing" @input="$emit('edit', $event, 'contact','phone')">{{ contact.phone }}</span>
        </li>
        <li>
            <img src="/email.png" class="icon">
            <span :contenteditable="editing" @input="$emit('edit', $event, 'contact','email')">{{ contact.email }}</span>
        </li>
        <li>
            <img src="/location.png" class="icon">
            <span :contenteditable="editing" @input="$emit('edit', $event, 'contact','address')">{{ contact.address }}</span>
        </li>
    </ul>
</template>

<script>
export default {
    props: {
        "contact": {
            type: Object
        },
        "editing": {
            type: Boolean,
            default: true
        },
    }
}
</script>

<style>

.icon {
    margin-right: 10px;
    margin-left: 0;
    vertical-align: middle;
    height: 18px;
    width: 18px;
    border-color: white;
}

.contact li {
    display: flex;
    align-items: center;
    flex-wrap: wrap;
    margin-bottom: 10px;
}

.contact span {
    width: 80%;
}

.contact li:last-child {
    margin-bottom: 0;
}

#resume .contact {
    padding-inline-start: 0;
}
</style>