/** @type {import('tailwindcss').Config} */
export default {
  content: ["./index.html","./src/**/*.{js,ts,jsx,tsx}"],
  theme: {
    extend: {
      colors: {
        black: "#303030",
        hint: "#707579",
        link: "#007AFF",
        accent: "#007AFF",
        purple: "#D05EDA",
        gold: "#FCC40C",
        divide: "#EFEFF4",
        success: "#00C067",
        btnText: "#FFFFFF",
        stopRed: "#FF1500",
        bezeled: "#007AFF26",
        inactive: "#A2ACB0",
        secondaryBg: "#EFEFF4",
        lightGray: "#8E8E92",
        darkGray: "#6D6D71",
      },
    },
  },
  plugins: [],
}

