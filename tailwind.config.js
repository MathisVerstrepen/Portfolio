/** @type {import('tailwindcss').Config} */
module.exports = {
    content: ["./components/**/*.{html,js,templ}"],
    theme: {
        extend: {
            colors: {
                base: 'rgb(var(--c_base)/<alpha-value>)',
                surface: 'rgb(var(--c_surface)/<alpha-value>)',
                gold: 'rgb(var(--c_gold)/<alpha-value>)',
                ink: 'rgb(var(--c_ink)/<alpha-value>)',
                smoke: 'rgb(var(--c_smoke)/<alpha-value>)',
                line: 'rgb(var(--c_line)/<alpha-value>)',
            },
            fontFamily: {
                sans: ["Manrope", "ui-sans-serif", "system-ui", "sans-serif"],
                serif: ["'Playfair Display'", "ui-serif", "Georgia", "serif"],
            },
            letterSpacing: {
                micro: "0.18em",
            },
            maxWidth: {
                prose: "68ch",
            },
        },
    },
    plugins: [],
};
