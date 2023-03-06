package service

import (
	"fmt"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
)

type Dashboard struct {
	service.Service
}

func (e *Dashboard) GetSalesByM(c *dto.DashboardReq, list *[]models.GetSalesByMRes, userId int) error {
	var err error
	sql := fmt.Sprint("SELECT " +
		"DATE_FORMAT(q.`created_at`,'%Y-%m') as sales_date," +
		"q.`create_by` as who," +
		"c.`name`," +
		"SUM(q.`grand_total`) as sales_sum " +
		"FROM `quote` q " +
		"LEFT JOIN `currency` c ON c.id = q.`currency_id` " +
		"where " +
		"q.`status` = 'Approved' AND " +
		"DATE_FORMAT(q.`created_at`,'%Y') = ? AND " +
		"c.`name` = ? AND " +
		"q.`create_by` != 0 and " +
		"q.`create_by` = ? " +
		"GROUP BY q.`create_by`, c.`name`, DATE_FORMAT(q.`created_at`,'%Y-%m') " +
		"ORDER BY q.`create_by`, DATE_FORMAT(q.`created_at`,'%Y-%m')")
	rows, err := e.Orm.Raw(sql, c.Year, c.Currency, userId).Rows()
	defer rows.Close()
	if err != nil {
		e.Log.Errorf("DashboardService GetSalesByM error:%s \r\n", err)
		return err
	}

	for rows.Next() {
		err = e.Orm.ScanRows(rows, list)
		if err != nil {
			return err
		}
	}

	return nil
}

func (e *Dashboard) GetSalesTop20(c *dto.DashboardReq, list *[]models.GetSalesTop20) error {
	var err error
	sql := fmt.Sprint("select " +
		"u.`nick_name` as name," +
		"SUM(q.`grand_total`) as sales_sum " +
		"from `quote` q  left join `currency` c on c.id = q.`currency_id` " +
		"left join `sys_user` u on u.`user_id` = q.`create_by` " +
		"where " +
		"q.`status` = 'Approved' and " +
		"DATE_FORMAT(q.`created_at`,'%Y') = ? and " +
		"q.`create_by` != 0 and " +
		"c.`name` = ? " +
		"group by q.`create_by`, c.`name`, DATE_FORMAT(q.`created_at`,'%Y') " +
		"order by sales_sum DESC limit 15")
	rows, err := e.Orm.Raw(sql, c.Year, c.Currency).Rows()
	defer rows.Close()
	if err != nil {
		e.Log.Errorf("DashboardService GetSalesTop20 error:%s \r\n", err)
		return err
	}

	for rows.Next() {
		err = e.Orm.ScanRows(rows, list)
		if err != nil {
			return err
		}
	}

	return nil
}

func (e *Dashboard) GetSalesByMAccount(c *dto.DashboardReq, list *[]models.GetSalesByMAccount, userId int) error {
	var err error
	sql := fmt.Sprint("select " +
		"A.sales_date, sum(A.account_number) as account_number from (" +
		"select DATE_FORMAT(q.`created_at`,'%Y-%m') as sales_date, if(count(1)>=1,1,0) as account_number " +
		"from quote q " +
		"LEFT JOIN `currency` c ON c.`id` = q.`currency_id` " +
		"where q.`create_by` = ? and " +
		"q.`create_by` != 0 and " +
		"DATE_FORMAT(q.`created_at`,'%Y') = ? and " +
		"q.`status` = 'Approved' and " +
		"c.`name` = ? " +
		"group by account_id, DATE_FORMAT(q.`created_at`,'%Y-%m') " +
		") A group by A.sales_date")
	rows, err := e.Orm.Raw(sql, userId, c.Year, c.Currency).Rows()
	defer rows.Close()
	if err != nil {
		e.Log.Errorf("DashboardService GetSalesByMAccount error:%s \r\n", err)
		return err
	}

	for rows.Next() {
		err = e.Orm.ScanRows(rows, list)
		if err != nil {
			return err
		}
	}

	return nil
}

func (e *Dashboard) GetSalesByProduct(c *dto.DashboardReq, list *[]models.GetSalesByProduct, userId int) error {
	var err error
	sql := fmt.Sprint("select " +
		"IFNULL(pt.name,'n/a') as product_name, " +
		"sum(qi.sales_price) as sales_sum " +
		"from `quote_item` qi " +
		"left join `products` p on p.`id` = qi.`product_id` " +
		"left join `products_type` pt on p.`type_id` = pt.`id` " +
		"left join `quote` q on q.`id` = qi.`quote_id` " +
		"left join `currency` c on c.`id` = qi.`currency_id` " +
		"where " +
		"c.`name` = ? and " +
		"date_format(q.`created_at`, '%Y') = ? and " +
		"q.`create_by` != 0 and " +
		"q.`create_by` = ? " +
		"group by pt.`name`, qi.`currency_id` " +
		"HAVING sales_sum > 0 " +
		"order by sales_sum desc limit 15")
	rows, err := e.Orm.Raw(sql, c.Currency, c.Year, userId).Rows()
	defer rows.Close()
	if err != nil {
		e.Log.Errorf("DashboardService GetSalesByProduct error:%s \r\n", err)
		return err
	}

	for rows.Next() {
		err = e.Orm.ScanRows(rows, list)
		if err != nil {
			return err
		}
	}

	return nil
}
