USE [compras_1]
GO
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE  PROC [dbo].[NS_TMMOTIVOS_EXTORNO_PEDIDO_I](
	@JSON VARCHAR(MAX)
)
AS
DECLARE @ErrorMessage NVARCHAR(max),
		@ErrorSeverity INT,
		@ErrorState INT,
		@id smallint;

DECLARE @temptable table(
	id smallint
)
BEGIN
	BEGIN TRY
	BEGIN TRANSACTION;

		WITH MOTIVOSEXTORNOPEDIDO (nombre,codigo,habilitado) AS (
			SELECT
				nombre,codigo,habilitado
			FROM OPENJSON(@JSON)
			WITH (
				nombre VARCHAR(250),codigo VARCHAR(15),habilitado BIT
			)
		)
		INSERT INTO TMMOTIVOS_EXTORNO_PEDIDO(
				nombre,codigo,habilitado
		)
		OUTPUT INSERTED.id INTO @temptable(id)
		SELECT nombre,codigo,habilitado
		FROM MOTIVOSEXTORNOPEDIDO

		SELECT @id = id FROM @temptable ;

		SELECT @id as id
		COMMIT TRANSACTION
	END TRY
	BEGIN CATCH
		ROLLBACK TRANSACTION;
		SELECT
			@ErrorMessage='Ocurrio un Error: '+ERROR_MESSAGE() + ' LN. ' + CONVERT(NVARCHAR(255), ERROR_LINE() ) + '.',
			@ErrorSeverity=ERROR_SEVERITY(),
			@ErrorState=ERROR_STATE();
		THROW 51003, @ErrorMessage, 10
	END CATCH

END
