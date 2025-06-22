import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { measurementCategories } from "@/constants/units";
import UnitConverterForm from "@/UnitConverterForm";

const App = () => (
  <div className="max-w-xl md:mx-auto mx-6">
    <h1 className="text-3xl font-bold my-10">Unit Converter</h1>

    <main>
      <UnitConverterTabs />
    </main>
  </div>
);

export default App;

const UnitConverterTabs = () => (
  <Tabs defaultValue="length">
    <TabsList>
      {measurementCategories.map((category) => (
        <TabsTrigger key={category} value={category} className="capitalize">
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
);
