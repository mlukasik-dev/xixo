export type Client = {
  id: string;
  name: string;
  owner: string;
  status: string;
  notes: string;
  email: string;
};

export const CLIENTS: Client[] = [
  {
    id: "12",
    name: "Annette Black",
    owner: "Wade Warren",
    status: "In Process",
    notes: "Add / Review",
    email: "serg@mail.ru",
  },
  {
    id: "34",
    name: "Guy Hawkins",
    owner: "Kristin Watson",
    status: "In Process",
    notes: "Add / Review",
    email: "serg@mail.ru",
  },
  {
    id: "22",
    name: "Cody Fisher",
    owner: "Brooklyn Simmons",
    status: "In Process",
    notes: "Add / Review",
    email: "serg@mail.ru",
  },
];
