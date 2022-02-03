import React from 'react';
import { makeStyles } from '@material-ui/styles';
import { Typography, TextField, Grid, Button, Box } from '@material-ui/core';
import { Paper } from '../../../../../components';

const useStyles = makeStyles((theme) => ({
  root: {
    marginTop: theme.spacing(3)
  },
  paper: { padding: theme.spacing(4) },
  gridContainer: {
    marginTop: theme.spacing(4)
  },
  successInfo: { margin: theme.spacing(3) },
  ignoreButton: {
    marginLeft: theme.spacing(3)
  }
}));

const convertToAlphabet = (value) => (value + 10).toString(36).toUpperCase();

export default function BookingInvitation(props) {
  const classes = useStyles(props);
  // const {
  //   selectedSeats,
  //   sendInvitations,
  //   ignore,
  //   invitations,
  //   onSetInvitation,
  //   onDownloadPDF
  // } = props;
  const { invitations, data } = props;

  const notValidInvitations = !Object.keys(invitations).length;

  return (
    <div className={classes.root}>
      <Paper className={classes.paper}>
        <Typography variant="h4" align="center">
          Check-in
        </Typography>
        <Typography
          className={classes.successInfo}
          variant="body1"
          align="center">
          You have successfuly booked your seats. Please scan QRCode when you
          arrive at cinema.
        </Typography>
        <Box
          sx={{
            display: 'flex',
            justifyContent: 'center',
            flexDirection: { xs: 'column', md: 'row' },
            alignItems: 'center',
            bgcolor: 'background.paper',
            overflow: 'hidden',
            borderRadius: '12px',
            boxShadow: 1,
            fontWeight: 'bold'
          }}>
          <Box
            component="img"
            // sx={{
            //   height: 233,
            //   width: 350,
            //   maxHeight: { xs: 233, md: 167 },
            //   maxWidth: { xs: 350, md: 250 }
            // }}
            alt="The house from the offer."
            src={data.QRCode}
          />
          <Box
            sx={{
              display: 'flex',
              flexDirection: 'column',
              alignItems: { xs: 'center', md: 'flex-start' },
              m: 3,
              minWidth: { md: 350 }
            }}>
            <Box component="span" sx={{ fontSize: 16, mt: 1 }}>
              {`Date: ${new Date(
                data.selectedDate
              ).toLocaleDateString()} - Time: ${data.selectedTime}`}
            </Box>
            <Box component="span" sx={{ color: 'primary.main', fontSize: 22 }}>
              {`Movie: ${data.movie.title} -
               Cinema: ${data.cinema.name}`}
            </Box>
          </Box>
        </Box>

        {/* <Box width={1} textAlign="center">
          <Button
            color="primary"
            variant="outlined"
            onClick={() => onDownloadPDF()}>
            Download Pass
          </Button>
          <Typography>

          </Typography>
        </Box> */}
        {/* <Grid className={classes.gridContainer} container spacing={3}>
          {selectedSeats.map((seat, index) => (
            <Grid item xs={12} md={6} lg={4} key={'seat-' + index}>
              <TextField
                fullWidth
                label="email"
                name={`${convertToAlphabet(seat[0])}-${seat[1]}`}
                helperText={`Please select an Email for Row : ${convertToAlphabet(
                  seat[0]
                )} - Seat Number : ${seat[1]}`}
                margin="dense"
                required
                value={
                  invitations[`${convertToAlphabet(seat[0])}-${seat[1]}`] || ''
                }
                variant="outlined"
                onChange={event => onSetInvitation(event)}
              />
            </Grid>
          ))}
          <Grid item xs={12} container>
            <Grid item>
              <Button
                disabled={notValidInvitations}
                color="primary"
                variant="outlined"
                onClick={() => sendInvitations()}>
                Send Invitations
              </Button>
            </Grid>
            <Grid item>
              <Button
                className={classes.ignoreButton}
                color="secondary"
                variant="outlined"
                onClick={() => ignore()}>
                Ignore
              </Button>
            </Grid>
          </Grid>
        </Grid> */}
      </Paper>
    </div>
  );
}
