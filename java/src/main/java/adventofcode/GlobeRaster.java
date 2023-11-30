package adventofcode;

import java.util.HashSet;
import java.util.Set;

public class GlobeRaster {


	public static void checkColissions(){
		Set<Long> values = new HashSet<>();
		GlobeRaster globeRaster = new GlobeRaster();
		double step = 1;
		for (double lat = -90.0; lat< 90.0; lat=lat+step){
			for (double lon = -180.0; lon< 180.0; lon=lon+step){
				long d1Coordinate = globeRaster.get1DCoordinateFromLatitudeAndLongitude(lat, lon);

				if (values.contains(d1Coordinate)){
					System.out.println(d1Coordinate);
				}else{
					values.add(d1Coordinate);
				}
//				if (values.size()%10000==0) {
//					System.out.println("d1Coordinate: " + d1Coordinate + "  lat: " + lat + "  lon " + lon);
//				}

				if (lat == 0 && lon ==0) {
					System.out.println("d1Coordinate: " + d1Coordinate + "  lat: " + lat + "  lon " + lon);
				}
			}
		}
		System.out.println(values.size());
	}
	public static void main(String[] args) {
    	checkColissions();
	}
	public static void main2(String[] args) {
		GlobeRaster globeRaster = new GlobeRaster();
		final double latitude = 78.02;
		final double longitude = -168.09;
		System.out.println("");
		System.out.println("Forward: (LAT/LNG via 1D COORDINATE to QUADRANT ID)");
		System.out.println("\tLatitude =<" + latitude + ">");
		System.out.println("\tLongitude =<" + longitude + ">");
		long d1Coordinate = globeRaster.get1DCoordinateFromLatitudeAndLongitude(latitude, longitude);
		System.out.println("\t1D-Coordinate =<" + d1Coordinate + ">");
		long quadrantId = globeRaster.getQuadrantIDFrom1DCoordinate(d1Coordinate);
		System.out.println("\tQuadrant ID =<" + quadrantId + ">");
		System.out.println("");
		System.out.println("Reverse: (QUADRANT ID via 1D-COORDINATE to CORNER LAT/LNG)");
		System.out.println("\tQuadrant ID =<" + quadrantId + ">");
		d1Coordinate = globeRaster.get1DCoordinateFromQuadrantID(quadrantId);
		System.out.println("\t1D-Coordinate =<" + d1Coordinate + ">");
		double[] latitudeLongitudeArray = globeRaster.getLatitudeAndLongitudeFrom1DCoordinates(d1Coordinate);
		System.out.println("\tLatidude =<" + latitudeLongitudeArray[0] + ">");
		System.out.println("\tLongitude =<" + latitudeLongitudeArray[1] + ">");
		System.out.println("");
		System.out.println("Short Forward:");
		System.out.println("\tLatitude =<" + latitude + ">");
		System.out.println("\tLongitude =<" + longitude + ">");
		quadrantId = globeRaster.getQuadrantIdFromLatitudeAndLongitude(latitude, longitude);
		System.out.println("\tQuadrant ID =<" + quadrantId + ">");
		System.out.println("");
		System.out.println("Short Reverse:");
		System.out.println("\tQuadrant ID =<" + quadrantId + ">");
		latitudeLongitudeArray = globeRaster.getLatitudeAndLongitudeFromQuadrantID(quadrantId);
		System.out.println("\tLatidude =<" + latitudeLongitudeArray[0] + ">");
		System.out.println("\tLongitude =<" + latitudeLongitudeArray[1] + ">");
		System.out.println("");
		System.out.println("Raster-Walk:");
		System.out.println("\tQuadrant ID =<" + quadrantId + ">");
		System.out.println("\tQuadrant ID Left =<" + globeRaster.getIdLeft(quadrantId) + ">");
		System.out.println("\tQuadrant ID Right =<" + globeRaster.getIdRight(quadrantId) + ">");
		System.out.println("\tQuadrant ID Above =<" + globeRaster.getIdAbove(quadrantId) + ">");
		System.out.println("\tQuadrant ID Beneath =<" + globeRaster.getIdBeneath(quadrantId) + ">");
		System.out.println("\tQuadrant ID Abobe Left =<" + globeRaster.getIdAboveLeft(quadrantId) + ">");
		System.out.println("\tQuadrant ID Abobe Right =<" + globeRaster.getIdAboveRight(quadrantId) + ">");
		System.out.println("\tQuadrant ID Beneath Left =<" + globeRaster.getIdBeneathLeft(quadrantId) + ">");
		System.out.println("\tQuadrant ID Beneath Right =<" + globeRaster.getIdBeneathRight(quadrantId) + ">");
	}
	public GlobeRaster() {
		precision = DEFAULT_PRECISION;
		setPrecision(precision);
	}
	public GlobeRaster(final double precision) {
		this.precision = precision;
		setPrecision(precision);
	}
	private void setPrecision(final double precision) {
		factor = (1.0d / precision);
		if (debug) {
			System.out.println("factor=<" + factor + ">");
		}
		setLatitudeRound();
	}
	private void setLatitudeRound() {
		latitudeRound = (int) (360 * factor);
		if (debug) {
			System.out.println("latitudeRound=<" + latitudeRound + ">");
		}
	}
	public long get1DCoordinateFromLatitudeAndLongitude(final double latitude, final double longitude) {
		double latitudeDelta = (latitude - 80.001d);
		double longitudeDelta = longitude + 180.001d;
		final long leftShift = (LEFTSHIFT) * (long) (factor);
		final int x = (int) (factor * latitudeDelta);
		final long leftPart = -x * leftShift;
		final long rightPart = (long) (factor * longitudeDelta);
		final long calculate1DCoordinate = leftPart + rightPart;
		if (debug) {
			System.out.println("latitude=<" + latitude + ">");
			System.out.println("longitude=<" + longitude + ">");
			System.out.println("precision=<" + precision + ">");
			System.out.println("latitudeDelta=<" + latitudeDelta + ">");
			System.out.println("longitudeDelta=<" + longitudeDelta + ">");
			System.out.println("x=<" + x + ">");
			System.out.println("leftPart=<" + leftPart + ">");
			System.out.println("rightPart=<" + rightPart + ">");
		}
		return calculate1DCoordinate;
	}
	public long getQuadrantIDFrom1DCoordinate(final long d1Coordinate) {
		final long leftPart = (long) (d1Coordinate / (LEFTSHIFT * factor));
		final long quadrantsPerRound = (long) (leftPart * 360 * factor);
		final long rowLeftPart = (long) (leftPart * (LEFTSHIFT * factor));
		final long rightPart = d1Coordinate - rowLeftPart;
		if (debug) {
			System.out.println("d1Coordinate=<" + d1Coordinate + ">");
			System.out.println("leftPart=<" + leftPart + ">");
			System.out.println("quadrantsPerRound=<" + quadrantsPerRound + ">");
			System.out.println("rightPart=<" + rightPart + ">");
		}
		return quadrantsPerRound + rightPart;
	}
	public long get1DCoordinateFromQuadrantID(final long id) {
		final long left = (long) (id / (360 * factor));
		final long IdleftPart = (long) (left * 360 * factor);
		final long idRightPart = id - IdleftPart;
		if (debug) {
			System.out.println("left=<" + left + ">");
			System.out.println("IdleftPart=<" + IdleftPart + ">");
			System.out.println("idRightPart=<" + idRightPart + ">");
		}
		return (long) (left * factor * LEFTSHIFT) + idRightPart;
	}
	public long getQuadrantIdFromLatitudeAndLongitude(final double latitude, final double longitude) {
		final long d1Coordinate = get1DCoordinateFromLatitudeAndLongitude(latitude, longitude);
		return getQuadrantIDFrom1DCoordinate(d1Coordinate);
	}
	public double[] getLatitudeAndLongitudeFromQuadrantID(final long quadrantId) {
		long d1Coordinate = get1DCoordinateFromQuadrantID(quadrantId);
		return getLatitudeAndLongitudeFrom1DCoordinates(d1Coordinate);
	}
	public double[] getLatitudeAndLongitudeFrom1DCoordinates(final long d1Coordinate) {
		// TODO
		System.err.println("getLatitudeAndLongitudeFrom1DCoordinates() --> HERE THE LOGIC IS MISSING");
		return new double[] { -1d, -1d };
	}
	public long getIdAbove(long id) {
		return (long) (id + (360 * factor));
	}
	public long getIdAboveLeft(long id) {
		return (long) (id + latitudeRound) - 1;
	}
	public long getIdAboveRight(long id) {
		return (long) (id + latitudeRound) + 1;
	}
	public long getIdLeft(long id) {
		return id - 1;
	}
	public long getIdRight(long id) {
		return id + 1;
	}
	public long getIdBeneath(long id) {
		return (long) (id - latitudeRound);
	}
	public long getIdBeneathLeft(long id) {
		return (long) (id - latitudeRound) - 1;
	}
	public long getIdBeneathRight(long id) {
		return (long) (id - latitudeRound) + 1;
	}
	public boolean isDebug() {
		return debug;
	}
	public void setDebug(boolean debug) {
		this.debug = debug;
	}
	private final static long LEFTSHIFT = 1000;
	private final static double DEFAULT_PRECISION = 0.01;
	private boolean debug = false;
	private double precision;
	private double factor;
	private int latitudeRound;
}



//	This code is a Java class called GlobeRaster, which appears to perform various operations related to converting latitude and longitude coordinates to a 1D coordinate and quadrant ID, and vice versa. The class also appears to include methods for "walking" the raster, such as getting the quadrant ID to the left, right, above, and beneath a given quadrant ID.
//
//		The main method of the class initializes an instance of GlobeRaster and sets the latitude and longitude to 78.02 and -168.09 respectively. It then performs various operations, such as:
//
//		get1DCoordinateFromLatitudeAndLongitude(latitude, longitude) which converts the given latitude and longitude to a 1D coordinate
//		getQuadrantIDFrom1DCoordinate(d1Coordinate) which converts the 1D coordinate to a quadrant ID
//		get1DCoordinateFromQuadrantID(quadrantId) which converts the quadrant ID back to a 1D coordinate
//		getLatitudeAndLongitudeFrom1DCoordinates(d1Coordinate) which converts the 1D coordinate back to a latitude and longitude
//		getQuadrantIdFromLatitudeAndLongitude(latitude, longitude) which is a short version of steps 1 and 2
//		getLatitudeAndLongitudeFromQuadrantID(quadrantId) which is a short version of steps 3 and 4
//		getIdLeft(quadrantId), getIdRight(quadrantId), getIdAbove(quadrantId), getIdBeneath(quadrantId), getIdAboveLeft(quadrantId), getIdAboveRight(quadrantId) which are used for raster-walking
//		It looks like the class does not include the implementation for these methods, therefore it would not execute correctly.
