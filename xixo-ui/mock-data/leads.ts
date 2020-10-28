export type Lead = {
  id: string;
  name: string;
  email: string;
  phone: string;
  date: string;
  type: string;
  source: string;
};

export const LEADS: Lead[] = [
  {
    id: "123456",
    name: "Martin Lukasik",
    email: "martilukas@gmail.com",
    phone: "+48 507 968 492",
    date: Date.now().toString(),
    type: "call",
    source: "Website",
  },
];
