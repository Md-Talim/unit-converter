import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { Button } from "./components/ui/button";
import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
  CardTitle,
} from "./components/ui/card";
import { Input } from "./components/ui/input";
import { Label } from "./components/ui/label";
import { measurementCategories, measurementUnits } from "./constants/units";

const App = () => (
  <div className="max-w-xl md:mx-auto mx-6">
    <h1 className="text-3xl font-bold my-10">Unit Converter</h1>

    <main className="space-y-6">
      <Tabs defaultValue="length">
        <TabsList>
          {measurementCategories.map((category) => (
            <TabsTrigger value={category} className="capitalize">
              {category}
            </TabsTrigger>
          ))}
        </TabsList>
        {measurementCategories.map((category) => (
          <TabsContent key={category} value={category}>
            <UnitConverterForm category={category} />
          </TabsContent>
        ))}
      </Tabs>
    </main>
  </div>
);

export default App;

interface FormProps {
  category: keyof typeof measurementUnits;
}

const UnitConverterForm = ({ category }: FormProps) => {
  const lengthUnits = measurementUnits[category];

  return (
    <Card>
      <CardHeader>
        <CardTitle className="capitalize">Convert {category}</CardTitle>
      </CardHeader>
      <CardContent>
        <form className="grid gap-6">
          <div className="space-y-1.5">
            <Label htmlFor="length">Enter the {category} to covert</Label>
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
