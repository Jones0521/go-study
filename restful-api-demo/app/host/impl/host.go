package impl

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/go-jones/restful-api-demo/app/host"
	"github.com/infraboard/mcube/sqlbuilder"
)

func (i *impl) CreateHost(ctx context.Context, ins *host.Host) (*host.Host, error) {
	// 校验数据合法性

	if err := ins.Validate(); err != nil {
		return nil, err
	}
	// 全局异常
	var (
		resStmt  *sql.Stmt
		descStmt *sql.Stmt
		err      error
	)
	// 把数据入库到resource表和host表
	// 一次需要往2个表录入数据, 我们需要二次插入操作，要么都成功,要么都失败, 引入事务逻辑
	// 初始化一个事务, 所有的的操作都使用这个事务来进行提交
	tx, err := i.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			err := tx.Rollback()
			i.log.Debugf("tx rollback error %s", err)
		} else {
			err := tx.Commit()
			if err != nil {
				i.log.Debugf("tx commit error, %s", err)
			}
		}
	}()
	// 在这个事务里面执行 Insert SQL, 先执行Prepare, 防止SQL注入
	resStmt, err = tx.Prepare(insertResourceSQL)
	if err != nil {
		return nil, fmt.Errorf("prepare resource sql error, %s", err)
	}
	defer resStmt.Close()
	// 注意: Prepare语句, 会占用Mysql 资源, 如果你使用后不关闭会导致Prepare溢出
	_, err = resStmt.Exec(
		ins.Id, ins.Vendor, ins.Region, ins.Zone, ins.CreateAt, ins.ExpireAt, ins.Category, ins.Type, ins.InstanceId,
		ins.Name, ins.Description, ins.Status, ins.UpdateAt, ins.SyncAt, ins.SyncAccount, ins.PublicIP,
		ins.PrivateIP, ins.PayType, ins.ResourceHash, ins.DescribeHash,
	)
	if err != nil {
		return nil, fmt.Errorf("insert resource error, %s", err)
	}
	// 同样的逻辑, 我们也需要Host的数据存入
	descStmt, err = tx.Prepare(insertDescribeSQL)
	if err != nil {
		return nil, fmt.Errorf("prepare describe sql error, %s", err)
	}
	defer descStmt.Close()
	_, err = descStmt.Exec(
		ins.Id, ins.CPU, ins.Memory, ins.GPUAmount, ins.GPUSpec, ins.OSType, ins.OSName,
		ins.SerialNumber, ins.ImageID, ins.InternetMaxBandwidthOut,
		ins.InternetMaxBandwidthIn, ins.KeyPairName, ins.SecurityGroups,
	)
	if err != nil {
		return nil, fmt.Errorf("insert describe error, %s", err)
	}

	return nil, nil
}

func (i *impl) QueryHost(ctx context.Context, req *host.QueryHostRequest) (*host.Set, error) {
	query := sqlbuilder.NewQuery(queryHostSQL).Order("create_at").Desc().Limit(int64(req.PageNumber), uint(req.PageSize))
	// build 查询语句
	sqlStr, args := query.BuildQuery()
	i.log.Debugf("sql: %s, args: %v", sqlStr, args)

	// Prepare
	stmt, err := i.db.Prepare(sqlStr)
	if err != nil {
		return nil, fmt.Errorf("prepare query host sql error, %s", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, fmt.Errorf("stmt query error, %s", err)
	}

	// 初始化需要返回的对象
	set := host.NewSet()

	// 迭代查询表里的数据
	for rows.Next() {
		ins := host.NewDefaultHost()
		if err := rows.Scan(
			&ins.Id, &ins.Vendor, &ins.Region, &ins.Zone, &ins.CreateAt, &ins.ExpireAt,
			&ins.Category, &ins.Type, &ins.InstanceId, &ins.Name,
			&ins.Description, &ins.Status, &ins.UpdateAt, &ins.SyncAt, &ins.SyncAccount,
			&ins.PublicIP, &ins.PrivateIP, &ins.PayType, &ins.ResourceHash, &ins.DescribeHash,
			&ins.Id, &ins.CPU,
			&ins.Memory, &ins.GPUAmount, &ins.GPUSpec, &ins.OSType, &ins.OSName,
			&ins.SerialNumber, &ins.ImageID, &ins.InternetMaxBandwidthOut, &ins.InternetMaxBandwidthIn,
			&ins.KeyPairName, &ins.SecurityGroups,
		); err != nil {
			return nil, err
		}

		set.Add(ins)
	}

	// Count 获取总数量
	// build 一个count语句
	countStr, countArgs := query.BuildCount()
	countStmt, err := i.db.Prepare(countStr)
	if err != nil {
		return nil, fmt.Errorf("prepare count stmt error, %s", err)
	}
	defer countStmt.Close()

	if err := countStmt.QueryRow(countArgs...).Scan(&set.Total); err != nil {
		return nil, fmt.Errorf("count stmt query error, %s", err)
	}

	return set, nil
}

func (i *impl) DescribeHost(ctx context.Context, req *host.DescribeHostRequest) (*host.Host, error) {
	return nil, nil
}

func (i *impl) UpdateHost(ctx context.Context, req *host.UpdateHostRequest) (*host.Host, error) {
	return nil, nil
}
func (i *impl) DeleteHost(ctx context.Context, req *host.DeleteHostRequest) (*host.Host, error) {
	return nil, nil
}
