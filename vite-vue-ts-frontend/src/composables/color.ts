export const hexToRgba = (hex: string, alphaOverride?: number) => {
  if (!hex) return `rgba(0,0,0,1)`;
  let h = hex.replace("#", "");
  let r,
    g,
    b,
    a = 1;
  if (h.length === 8) {
    // RRGGBBAA
    r = parseInt(h.slice(0, 2), 16);
    g = parseInt(h.slice(2, 4), 16);
    b = parseInt(h.slice(4, 6), 16);
    a = parseInt(h.slice(6, 8), 16) / 255;
  } else {
    // RRGGBB
    r = parseInt(h.slice(0, 2), 16);
    g = parseInt(h.slice(2, 4), 16);
    b = parseInt(h.slice(4, 6), 16);
  }

  const alpha = alphaOverride ?? a;

  return `rgba(${r}, ${g}, ${b}, ${alpha})`;
};

export const tagColor = (base: string) => {
  console.log(base);
  console.log(hexToRgba(base, 1));
  console.log({
    color: hexToRgba(base, 0.2),
    textColor: hexToRgba(base, 1),
    borderColor: hexToRgba(base, 0.5),
  });
  return {
    color: hexToRgba(base, 0.2),
    textColor: hexToRgba(base, 1),
    borderColor: hexToRgba(base, 0.5),
  };
};

export const getNaiveUITagColorProperty = (base: string) => {
  return {
    color: hexToRgba(base, 0.2),
    textColor: hexToRgba(base, 1),
    borderColor: hexToRgba(base, 0.5),
  };
};
