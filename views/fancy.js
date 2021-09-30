const sectionsId = [
	"#section0", "#section1", 
	"#section2", "#section3", 
	"#section4", "#section5"
];

gsap.registerPlugin(ScrollTrigger);

function oddObject(id) {
	return {
		scrollTrigger: {
			trigger: id,
			toggleActions: "restart none none none"
		},
		duration: 0.25,
		opacity: 0,
		x: "-15px",
		delay: 0.5
	}
}

function evenObject(id) {
	return {
		scrollTrigger: {
			trigger: id,
			toggleActions: "restart none none none"
		},
		duration: 0.25,
		opacity: 0,
		x: "15px",
		delay: 0.5
	}
}

for(let i = 0; i < sectionsId.length; i++) {
	let object = evenObject(sectionsId[i]);
	if(i % 3 == 0) object = oddObject(sectionsId[i]);
	gsap.from(sectionsId[i], object);
}