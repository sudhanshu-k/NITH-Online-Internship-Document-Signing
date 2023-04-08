import React from "react";
import { AppBar, Toolbar, Typography, Drawer, List, ListItem, Divider, ListItemButton, Box } from "@mui/material";
import { ListItemIcon, ListItemText } from "@mui/material";
import { Mail as MailIcon } from "@mui/icons-material";
import "./Sidebar.css";

function Sidebar() {
	return (
		<div>
			<Drawer
				variant="permanent"
				sx={{
					width: 260,
					flexShrink: 0,
					boxSizing: "border-box",
					[`& .MuiDrawer-paper`]: { width: 260, boxSizing: "border-box" },
				}}
			>
				<Toolbar disableGutters="true"/>
				<Box sx={{ overflow: "auto" }}>
					<List>
						<ListItem disablePadding>
							<ListItemButton>
								<ListItemIcon>
									<MailIcon />
								</ListItemIcon>
								<ListItemText>List Item 1</ListItemText>
							</ListItemButton>
						</ListItem>
					</List>
					<Divider />
					<List>
						<ListItem disablePadding>
							<ListItemButton>
								<ListItemIcon>
									<MailIcon />
								</ListItemIcon>
								<ListItemText>List Item 2</ListItemText>
							</ListItemButton>
						</ListItem>
					</List>
				</Box>
			</Drawer>
		</div>
	);
}

export default Sidebar;
