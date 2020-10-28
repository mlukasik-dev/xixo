import { FC } from "react";
import { Typography } from "@material-ui/core";
import { Meta, Story } from "@storybook/react/types-6-0";

const Title: FC = ({ children }) => <Typography>{children}</Typography>;

export default {
  title: "UI/Title",
  component: Title,
} as Meta;

const Template: Story = (args) => <Title {...args} />;

export const Default = Template.bind({});
Default.args = {
  children: "Title",
};
