//kage:unit pixels

//go:build ignore

package game

// Uniform variables.
var LightPosition vec2 // The position of the light source
var LightRadius float  // The radius of the light

// Fragment is the entry point of the fragment shader.
func Fragment(dstPos vec4, srcPos vec2, color vec4) vec4 {
	// Calculate the distance from the light source
	lightDist := length(LightPosition - dstPos.xy)

	// Calculate the light factor, where 0.0 is full light, and 1.0 is full darkness
	lightFactor := smoothstep(LightRadius, LightRadius*0.8, lightDist)

	// Apply the light factor to the original color, preserving the original alpha
	return vec4(color.rgb*(1.0-lightFactor), color.a)
}
