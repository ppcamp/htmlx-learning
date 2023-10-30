/** @type {import('tailwindcss').Config} */
module.exports = {
    // if you use content, it'll miss some tailwind classes (use "purge" instead)
    content: ["./templates/**/*.{html,tmpl}",],
    mode: "jit",
    theme: {
        extend: {},
    },
    plugins: [],
}

