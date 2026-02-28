package ui

func GetStyle(v Style, component ...string) string {
	// component peut être "btn", "badge", "alert"
	prefix := ""
	if len(component) > 0 && component[0] != "" {
		prefix = component[0] + "-"
	}
	switch v {
	case StyleOutline:
		return prefix + "-outline border-base-300 bg-transparent" // ex: "btn-outline"
	case StyleGhost:
		return prefix + "-ghost bg-base-200/50 border-none"
	case StyleGlass:
		return "glass backdrop-blur-md"
	default:
		return "bg-base-200 border border-base-200"
	}
}

func GetAvatarSize(s Size) string {
	switch s {
	case SizeXS:
		return "w-8"
	case SizeSM:
		return "w-12"
	case SizeMD:
		return "w-16"
	case SizeLG:
		return "w-24"
	case SizeXL:
		return "w-32"
	default:
		return "w-16"
	}
}

func GetPadding(s Size) string {
	switch s {
	case SizeSM:
		return "p-2 md:p-4" // Padding responsive automatique !
	case SizeMD:
		return "p-4 md:p-6"
	case SizeLG:
		return "p-6 md:p-10"
	default:
		return "p-4"
	}
}

func GetStateClasses(isLoading bool, isDisabled bool, hasError bool) string {
	var classes []string

	if isLoading {
		classes = append(classes, "loading loading-spinner opacity-70 cursor-not-allowed")
	}
	if isDisabled {
		classes = append(classes, "opacity-50 pointer-events-none grayscale")
	}
	if hasError {
		classes = append(classes, "border-error text-error")
	}

	return TwMerge(classes...)
}

func GetPosition(component string, pos Placement) string {
	// Si composant = "tooltip" et pos = PositionTop -> "tooltip-top"
	if pos == "" {
		return ""
	}
	return component + "-" + string(pos)
}

func GetContainer(mw MaxWidth) string {
	base := "container mx-auto px-4 sm:px-6 lg:px-8 w-full"
	if mw != "" {
		return base + " " + string(mw)
	}
	return base + " max-w-7xl" // Valeur par défaut robuste
}

func GetGrid(colCount int, gap Size) string {
	baseGrid := "grid grid-cols-1 w-full"

	// On gère l'espacement via notre helper de taille existant
	gapClass := "gap-4" // défaut
	switch gap {
	case SizeSM:
		gapClass = "gap-2 sm:gap-4"
	case SizeMD:
		gapClass = "gap-4 md:gap-6"
	case SizeLG:
		gapClass = "gap-6 lg:gap-10"
	}

	// On gère les colonnes Desktop
	colsClass := ""
	switch colCount {
	case 2:
		colsClass = "sm:grid-cols-2"
	case 3:
		colsClass = "md:grid-cols-2 lg:grid-cols-3"
	case 4:
		colsClass = "sm:grid-cols-2 lg:grid-cols-4"
	}

	return TwMerge(baseGrid, colsClass, gapClass)
}

func GetImageRatio(r Ratio, isCover bool) string {
	classes := string(r)
	if isCover {
		classes += " object-cover w-full h-full"
	}
	return classes
}

func GetProse() string {
	// Limite la largeur du texte, améliore l'interligne et la couleur
	return "max-w-prose mx-auto leading-relaxed text-base-content/80 text-lg"
}
