import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { measurementUnits, type MeasurementCategory } from "@/constants/units";
import { useState } from "react";

interface UnitConverterFormProps {
  category: MeasurementCategory;
}

const UnitConverterForm = ({ category }: UnitConverterFormProps) => {
  const lengthUnits = measurementUnits[category];

  const [value, setValue] = useState(0);
  const [selectedFromUnit, setSelectedFromUnit] = useState(lengthUnits[0]);
  const [selectedToUnit, setSelectedToUnit] = useState(lengthUnits[1]);

  const handleValueChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setValue(Number(e.target.value));
  };

  const handleFromUnitChange = (unit: string) => {
    setSelectedFromUnit(unit);
  };

  const handleToUnitChange = (unit: string) => {
    setSelectedToUnit(unit);
  };

  const handleConvert = async (
    e: React.FormEvent<HTMLFormElement | HTMLButtonElement>
  ) => {
    e.preventDefault();
    const params = new URLSearchParams({
      value: value.toString(),
      from: selectedFromUnit.toLowerCase(),
      to: selectedToUnit.toLowerCase(),
    });
    const url = `http://localhost:4000/convert?${params.toString()}`;

    console.log(url);

    const response = await fetch(url, { method: "POST" });
    if (!response.ok) {
      alert("Conversion failed");
      return;
    }

    const data = await response.json();
    alert(`Result: ${data.result}`);
  };

  return (
    <Card>
      <CardHeader>
        <CardTitle className="capitalize">Convert {category}</CardTitle>
      </CardHeader>
      <CardContent>
        <form className="grid gap-6" onSubmit={handleConvert}>
          <div className="space-y-1.5">
            <Label htmlFor="length">Enter the {category} to convert</Label>
            <Input
              id="length"
              name="length"
              type="number"
              value={value}
              onChange={handleValueChange}
            />
          </div>

          <div className="space-y-1.5">
            <Label htmlFor="from" className="">
              Unit to convert from
            </Label>
            <Select name="from" onValueChange={handleFromUnitChange}>
              <SelectTrigger className="w-full">
                <SelectValue placeholder="Convert from" />
              </SelectTrigger>
              <SelectContent>
                {lengthUnits.map((unit) => (
                  <SelectItem key={unit} value={unit}>
                    {unit}
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>
          </div>

          <div className="space-y-1.5">
            <Label htmlFor="to" className="">
              Unit to convert to
            </Label>
            <Select name="to" onValueChange={handleToUnitChange}>
              <SelectTrigger className="w-full">
                <SelectValue placeholder="Convert from" />
              </SelectTrigger>
              <SelectContent>
                {lengthUnits.map((unit) => (
                  <SelectItem key={unit} value={unit}>
                    {unit}
                  </SelectItem>
                ))}
              </SelectContent>
            </Select>
          </div>
        </form>
      </CardContent>
      <CardFooter>
        <Button
          type="button"
          className="hover:cursor-pointer"
          onClick={handleConvert}
        >
          Convert
        </Button>
      </CardFooter>
    </Card>
  );
};

export default UnitConverterForm;
