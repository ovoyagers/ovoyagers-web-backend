package dashboardservice

func (d *DashboardService) GetWebAnalytics() (map[string]interface{}, error) {
	return d.dashboardDao.GetWebAnalytics()
}