import type { Component } from 'solid-js';

const Nav: Component = () => {
  return (
    <nav class="z-10 top-0 sticky bg-white w-screen flex justify-center items-center py-2.5 md:py-4 shadow-lg shadow-slate-200">
      <div class="text-slate-600 text-3xl md:text-[40px] tracking-wide">
        moodboard
      </div>
      <div class="absolute right-5 md:right-8 flex items-center">
        <a target="_blank" href="https://felicepng.com">
          <img
            src="/src/assets/portfolio.png"
            alt="My Portfolio"
            class="-mt-0.5 h-[26px] md:h-8 w-[26px] md:w-8"
          />
        </a>
      </div>
    </nav>
  );
};

export default Nav;
