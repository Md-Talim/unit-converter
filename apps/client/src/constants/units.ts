export const measurementUnits = {
  length: [
    "Millimeters",
    "Centimeters",
    "Meters",
    "Kilometers",
    "Inches",
    "Feet",
    "Yards",
    "Miles",
  ],
  weight: ["Milligrams", "Grams", "Kilograms", "Pounds", "Ounces"],
  temperature: ["Celsius", "Fahrenheit", "Kelvin"],
};

export type MeasurementCategory = keyof typeof measurementUnits;

export const measurementCategories: MeasurementCategory[] = Object.keys(
  measurementUnits
) as MeasurementCategory[];
