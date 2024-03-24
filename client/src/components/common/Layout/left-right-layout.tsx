import { Layout, Menu, theme } from "antd";
import {
  DesktopOutlined,
  PieChartOutlined,
  FileOutlined,
} from "@ant-design/icons";
import Link from "~/assets/link.png";
import "./left-right-layout.scss";

const { Header, Content, Sider } = Layout;

function LeftRightLayout() {
  const {
    token: { colorBgContainer, borderRadiusLG },
  } = theme.useToken();
  return (
    <Layout style={{ minHeight: "100vh" }}>
      <Header style={{ padding: 0 }} className="headerContainer">
        <div className="logoContainer ">
          <img
            src={Link}
            alt="link"
            style={{ width: "25px", height: "25px", marginLeft: "25px" }}
          />
        </div>
        <div className="headerAccountContainer">Account</div>
      </Header>

      <Layout>
        <Sider>
          <Menu defaultSelectedKeys={["1"]} mode="inline">
            <Menu.Item key="1" icon={<PieChartOutlined />}>
              Option 1
            </Menu.Item>
            <Menu.Item key="2" icon={<DesktopOutlined />}>
              Option 2
            </Menu.Item>
            <Menu.Item key="3" icon={<FileOutlined />}>
              Option 3
            </Menu.Item>
          </Menu>
        </Sider>
        <Layout className="site-layout">
          <Content style={{ margin: "10px 16px" }}>
            <Content
              style={{
                padding: "0 24px",
                minHeight: 280,
                backgroundColor: colorBgContainer,
                borderRadius: borderRadiusLG,
              }}
            >
              Content
            </Content>
          </Content>
        </Layout>
      </Layout>
    </Layout>
  );
}

export default LeftRightLayout;
