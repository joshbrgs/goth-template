
// Create bubbles
const bubbles = document.querySelectorAll('.bubble');

// Animate each bubble
bubbles.forEach((bubble, index) => {
  gsap.to(bubble, {
    delay: index * 0.5, // Delay each bubble animation
    y: "-200%", // Move the bubble upwards
    opacity: 1, // Fade in the bubble
    duration: 2, // Animation duration
    ease: "power4.out", // Easing function
    repeat: -1, // Repeat animation infinitely
    yoyo: true, // Reverse animation
  });
});
