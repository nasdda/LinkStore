import { Menu } from "antd";

function NavBar() {
  return (
    <Menu mode="horizontal">
      <Menu.Item key="home">Home</Menu.Item>
      <Menu.Item key="about">About</Menu.Item>
      <Menu.Item key="contact">Contact</Menu.Item>
    </Menu>
  );
}

export default NavBar;
