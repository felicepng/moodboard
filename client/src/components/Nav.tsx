import type { Component } from 'solid-js';

const Nav: Component = () => {
  return (
    <nav class="z-10 top-0 sticky bg-slate-50 w-screen flex justify-center items-center py-2.5 sm:py-3.5 border-b border-slate-400">
      <div class="text-slate-900 text-3xl sm:text-[40px] tracking-wide">
        moodboard
      </div>
      <div class="absolute right-5 sm:right-8 flex items-center">
        <a target="_blank" href="https://felicepng.com">
          <img
            src="/src/assets/portfolio.png"
            alt="My Portfolio"
            class="-mt-0.5 h-[26px] sm:h-8 w-[26px] sm:w-8"
          />
        </a>
      </div>
    </nav>
  );
};

export default Nav;
