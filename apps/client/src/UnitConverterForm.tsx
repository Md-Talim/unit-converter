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

interface UnitConverterFormProps {
  category: MeasurementCategory;
}

const UnitConverterForm = ({ category }: UnitConverterFormProps) => {
  const lengthUnits = measurementUnits[category];

  return (
    <Card>
      <CardHeader>
        <CardTitle className="capitalize">Convert {category}</CardTitle>
      </CardHeader>
      <CardContent>
        <form className="grid gap-6">
          <div className="space-y-1.5">
            <Label htmlFor="length">Enter the {category} to convert</Label>
            <Input id="length" name="length" type="number" />
          </div>

          <div className="space-y-1.5">
            <Label htmlFor="from" className="">
              Unit to convert from
            </Label>
            <Select name="from">
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
            <Select name="to">
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
        <Button className="hover:cursor-pointer">Convert</Button>
      </CardFooter>
    </Card>
  );
};

export default UnitConverterForm;
