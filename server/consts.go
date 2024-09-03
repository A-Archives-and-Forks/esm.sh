package server

const (
	// esm.sh build version
	VERSION = 135
	// esm.sh stable build version, used for UI libraries like react, to make sure the runtime is single copy
	// change this carefully!
	STABLE_VERSION = 135
)

const (
	nodejsMinVersion = 20
	nodejsLatestLTS  = "20.11.1"
	nodeTypesVersion = "20.11.20"
	denoStdVersion   = "0.177.1"
)

// fix some npm package versions
var fixedPkgVersions = map[string]string{
	"@types/react@17": "17.0.71",
	"@types/react@18": "18.2.38",
	"isomorphic-ws@4": "5.0.0",
	"resolve@1.22":    "1.22.2", // 1.22.3+ will read package.json from disk
}

// css packages
var cssPackages = map[string]string{
	"@unocss/reset":    "tailwind.css",
	"inter-ui":         "inter.css",
	"normalize.css":    "normalize.css",
	"modern-normalize": "modern-normalize.css",
	"reset-css":        "reset.css",
}

// stable build for UI libraries like react, to make sure the runtime is single copy
var stableBuild = map[string]bool{
	"preact":            true,
	"react":             true,
	"solid-js":          true,
	"svelte":            true,
	"vue":               true,
	"@vue/reactivity":   true,
	"@vue/runtime-core": true,
	"@vue/runtime-dom":  true,
	"@vue/shared":       true,
}

var assetExts = map[string]bool{
	"wasm":       true,
	"css":        true,
	"less":       true,
	"sass":       true,
	"scss":       true,
	"stylus":     true,
	"styl":       true,
	"json":       true,
	"jsonc":      true,
	"csv":        true,
	"xml":        true,
	"plist":      true,
	"tmLanguage": true,
	"tmTheme":    true,
	"yml":        true,
	"yaml":       true,
	"pdf":        true,
	"txt":        true,
	"glsl":       true,
	"frag":       true,
	"vert":       true,
	"md":         true,
	"mdx":        true,
	"markdown":   true,
	"html":       true,
	"htm":        true,
	"vue":        true,
	"svelte":     true,
	"svg":        true,
	"png":        true,
	"jpg":        true,
	"jpeg":       true,
	"webp":       true,
	"gif":        true,
	"ico":        true,
	"eot":        true,
	"ttf":        true,
	"otf":        true,
	"woff":       true,
	"woff2":      true,
	"m4a":        true,
	"mp3":        true,
	"m3a":        true,
	"ogg":        true,
	"oga":        true,
	"wav":        true,
	"weba":       true,
	"mp4":        true,
	"m4v":        true,
	"ogv":        true,
	"webm":       true,
	"zip":        true,
	"gz":         true,
	"tar":        true,
	"tgz":        true,
}

// native node packages, for `denonext` target use `npm:package` instead of url
var nativeNodePackages = []string{
	"@achingbrain/ssdp",
	"default-gateway",
	"fsevent",
	"lightningcss",
	"re2",
	"zlib-sync",
}

var denoNextUnspportedNodeModules = map[string]bool{
	"inspector": true,
}
