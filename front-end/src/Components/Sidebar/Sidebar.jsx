import React from "react";
import { AppBar, Toolbar, Typography, Drawer, List, ListItem, Divider, ListItemButton, Box, Badge } from "@mui/material";
import { ListItemIcon, ListItemText } from "@mui/material";
import { Mail as MailIcon } from "@mui/icons-material";
import { useNavigate } from "react-router-dom";

function Sidebar() {
	const navigate = useNavigate();
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
				<Toolbar disableGutters={true} />
				<Box sx={{ overflow: "auto" }}>
					<List>
						<ListItem>
							<ListItemButton>
								<ListItemIcon>
									<Badge color="warning" variant="dot">
										<MailIcon />
									</Badge>
								</ListItemIcon>
								<ListItemText>Pending</ListItemText>
							</ListItemButton>
						</ListItem>
						<Divider />
						<ListItem>
							<ListItemButton>
								<ListItemIcon>
									<Badge color="success" variant="dot">
										<MailIcon />
									</Badge>
								</ListItemIcon>
								<ListItemText>Approved</ListItemText>
							</ListItemButton>
						</ListItem>
						<Divider />
						<ListItem>
							<ListItemButton>
								<ListItemIcon>
									<Badge color="error" variant="dot">
										<MailIcon />
									</Badge>
								</ListItemIcon>
								<ListItemText>Rejected</ListItemText>
							</ListItemButton>
						</ListItem>
					</List>
				</Box>
			</Drawer>
		</div>
	);
}

export default Sidebar;
